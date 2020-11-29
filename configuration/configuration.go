package configuration

import (
	"github.com/tkanos/gonfig"
	"log"
)

// Configuration contains configuration fields
type Configuration struct {
	HostPort string
}

// GetConfig returns config file
func GetConfig() Configuration {
	configuration := Configuration{}
	err := gonfig.GetConf("./configuration/config.json", &configuration)
	if err != nil {
		log.Printf("Config %s file not found", "configuration.json")
	}

	return configuration
}