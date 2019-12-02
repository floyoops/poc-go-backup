package main

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type ParametersDb struct {
	Parameters struct {
		DbDriver   string `yaml:"database_driver"`
		DbHost     string `yaml:"database_host"`
		DbPort     string `yaml:"database_port"`
		DbName     string `yaml:"database_name"`
		DbUser     string `yaml:"database_user"`
		DbPassword string `yaml:"database_password"`
	}
}

func YmlToParameters(byteValue []byte) (ParametersDb, error) {

	// Unmarshal
	p := ParametersDb{}
	err := yaml.Unmarshal([]byte(byteValue), &p)
	if err != nil {
		log.Fatalf("error: %v", err)
		os.Exit(1)
	}

	return p, err
}
