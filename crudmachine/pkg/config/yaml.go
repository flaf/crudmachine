package config

import(
	"log"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type Actions struct {
	Create string
	Read string
	Update string
	Delete string
}

type Conf struct {
	Server struct {
		Address string
		Port int
		Subject string
		Queue string
		Token string
	}
	Actions Actions
}

func (c *Conf) GetConf(conffile string) *Conf {

	yamlFile, err := ioutil.ReadFile(conffile)
	if err != nil {
	log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

