package internals

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type configurations struct {
	Kafka configKafka
}

type configKafka struct {
	Host      string
	Partition int
	MinBytes  int
	MaxBytes  int
}

func initConfigs() (c configurations) {
	config := configurations{}

	path, _ := os.Getwd()

	yamlFile, err := ioutil.ReadFile(path + "/resources/application.yml")
	if err != nil {
		log.Fatalf("Error to get application.yml properties, details: %v", err)
	}

	err = yaml.Unmarshal([]byte(yamlFile), &config)
	if err != nil {
		log.Fatalf("Error to unmarshal: %v", err)
	}
	return config
}
