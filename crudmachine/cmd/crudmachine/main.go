package main

import(
	"log"
	"os"
	"flag"
	"strconv"
	"ac-versailles/crudmachine/pkg/config"
	"ac-versailles/crudmachine/pkg/nats"
)

func usage() {
	log.Printf("Usage: crudmachine [-config <file>]\n")
	flag.PrintDefaults()
}

func showUsageAndExit(exitcode int) {
	usage()
	os.Exit(exitcode)
}

func main() {
	var configFile = flag.String("config", "/etc/crudmachine/config.yml", "The YAML configuration file")
	var showHelp = flag.Bool("help", false, "Show help message")
	log.SetFlags(0)
	flag.Usage = usage
	flag.Parse()

	if *showHelp {
		showUsageAndExit(0)
	}

	var c config.Conf
	c.GetConf(*configFile)

	nats.Listen(c.Server.Address + ":" + strconv.Itoa(c.Server.Port),
		c.Server.Subject, c.Server.Queue, c.Server.Token, c.Actions)

}

