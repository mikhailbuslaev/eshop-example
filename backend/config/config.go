package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func Parse(any interface{}, configFile string) {
	buf, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatalln("Cannot read config file")
	}

	err = yaml.Unmarshal(buf, any)
	if err != nil {
		log.Fatalln("Cannot parse config")
	}
}
