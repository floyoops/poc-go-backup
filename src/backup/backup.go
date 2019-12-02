package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// get args.
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Please enter a yml file path")
		os.Exit(1)
	}
	pathFile := args[1]

	// Open.
	byteValue, err := GetContent(pathFile)
	if err != nil {
		log.Fatalf("error on open file: %v\n", err)
		os.Exit(1)
	}

	// Unmarshal
	p, err := YmlToParameters(byteValue)
	if err != nil {
		log.Fatal("error on yml to parameters")
		os.Exit(1)
	}
	fmt.Printf("unmarshall %v success\n", pathFile)
	fmt.Printf("dbs %v", p.Parameters)
}
