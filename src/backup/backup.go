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

func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func main() {
	// get args.
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Please enter a yml file path")
		os.Exit(1)
	}
	pathFile := args[1]

	// Check if file exist.
	if Exists(pathFile) == false {
		fmt.Printf("the file '%v' does not exist\n", pathFile)
		os.Exit(1)
	}
	fmt.Printf("Path file yml detected: %v\n", pathFile)

	// Open.
	file, err := os.Open(pathFile)
	if err != nil {
		log.Fatalf("error on open file: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("open %v success\n", pathFile)

	// Read
	defer file.Close()
	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("error on read %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("read %v success\n", pathFile)

	// Unmarshal
	p := ParametersDb{}
	err = yaml.Unmarshal([]byte(byteValue), &p)
	if err != nil {
		log.Fatalf("error: %v", err)
		os.Exit(1)
	}
	fmt.Printf("unmarshall %v success\n", pathFile)

	fmt.Printf("dbs %v", p.Parameters)
}
