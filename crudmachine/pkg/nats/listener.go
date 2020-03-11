package nats

import (
	"runtime"
	"log"
//	"math/rand"
	"text/template"
	"bytes"
	"time"
	"encoding/json"
	"ac-versailles/crudmachine/pkg/exec"
	"ac-versailles/crudmachine/pkg/config"
	"github.com/nats-io/nats.go"
	"github.com/spf13/cast"
)

func handleMsg(m *nats.Msg, i int, actions config.Actions, ch chan int) {
	log.Printf("[#%d] Received on [%s]: '%s'", i, m.Subject, string(m.Data))

	//rand.Seed(time.Now().UnixNano())
	//min := 2
	//max := 4
	//rand_n := rand.Intn(max - min + 1) + min
	//time.Sleep(time.Duration(rand_n) * time.Second)

	t1 := time.Now()

	var dat map[string]interface{}
	if err := json.Unmarshal(m.Data, &dat); err != nil {
		panic(err)
	}


	action := cast.ToString(dat["action"])
	//action := dat["action"]
	//if action == nil {
	//	action = ""
	//} else {
	//	action = action.(string)
	//}
	data := dat["data"]

	if data == nil {
		log.Printf("data is nil")
	}

	log.Printf("[#%d] Action is: %s", i, action)

    cmd := ""
	switch action {
		case "create":
			cmd = actions.Create
		case "read":
			cmd = actions.Read
		case "update":
			cmd = actions.Update
		case "delete":
			cmd = actions.Delete
		default:
			log.Printf("[#%d] action [%s] not supported", i, action)
			return
	}

	//tmpl, err := template.New("cmdtpl").Option("missingkey=error").Parse(cmd)
	tmpl, err := template.New("cmdtpl").Parse(cmd)
	if err != nil { panic(err) }
	var b bytes.Buffer
	err = tmpl.Execute(&b, data)
	if err != nil { panic(err) }
	
	cmd = b.String()

	exec.Run(i, "/bin/sh", "-c", cmd)

	t2 := time.Now()
	diff := t2.Sub(t1)

	log.Printf("[#%d] Message handled after %d seconds", i, int(diff.Seconds()))

	// to remove one value from the channel.
	<-ch
}

func setupConnOptions(opts []nats.Option) []nats.Option {
	totalWait := 10 * time.Minute
	reconnectDelay := time.Second

	opts = append(opts, nats.ReconnectWait(reconnectDelay))
	opts = append(opts, nats.MaxReconnects(int(totalWait/reconnectDelay)))
	opts = append(opts, nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
		log.Printf("Disconnected due to:%s, will attempt reconnects for %.0fm", err, totalWait.Minutes())
	}))
	opts = append(opts, nats.ReconnectHandler(func(nc *nats.Conn) {
		log.Printf("Reconnected [%s]", nc.ConnectedUrl())
	}))
	opts = append(opts, nats.ClosedHandler(func(nc *nats.Conn) {
		log.Fatalf("Exiting: %v", nc.LastError())
	}))
	return opts
}

func Listen (server, subj, queue, token string, actions config.Actions) {
	opts := []nats.Option{nats.Name("NATS-crudmachine"), nats.Token(token)}

	nc, err := nats.Connect(server, opts...)
	if err != nil {
		log.Fatal(err)
	}

	i := 0

	// Number of concurrent tasks.
	ch := make(chan int, 10)

	nc.QueueSubscribe(subj, queue, func(msg *nats.Msg) {
		i += 1

		select {
			case ch <- 1:
				go handleMsg(msg, i, actions, ch)
			default:
				log.Printf("[#%d] Received on [%s] but skipped because stack is full: '%s'", i, msg.Subject, string(msg.Data))
		}

	})
	nc.Flush()

	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Listening on [%s]", subj)
	log.SetFlags(log.LstdFlags)

	runtime.Goexit()

}

