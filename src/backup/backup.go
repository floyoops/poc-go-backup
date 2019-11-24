package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
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

func main() {
	// Open
	pathYmlFile := "data/parameters.yml"
	file, err := os.Open(pathYmlFile)
	if err != nil {
		log.Fatalf("error on open file: %v\n", err)
	}
	fmt.Printf("open %v success\n", pathYmlFile)

	// Read
	defer file.Close()
	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("error on read %s\n", err)
	}
	fmt.Printf("read %v success\n", pathYmlFile)

	// Unmarshal
	p := ParametersDb{}
	err = yaml.Unmarshal([]byte(byteValue), &p)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("unmarshall %v success\n", pathYmlFile)

	fmt.Printf("dbs %v", p.Parameters)
}
