package main

import (
	app "backupYmlToFtp/src/backup/app/query"
	domain2 "backupYmlToFtp/src/backup/domain/adapter"
	domain "backupYmlToFtp/src/backup/domain/repository"
	"backupYmlToFtp/src/backup/infra/adapter"
	infra "backupYmlToFtp/src/backup/infra/repository"
	"fmt"
	"log"
	"os"
)

func main() {
	// bootstrap handler
	var fileRepo domain.FileRepository = infra.FileRepository{}
	var converter domain2.ConvertToParametersDb = adapter.YmlToParameters{}
	var handler = app.GetParametersDbQueryHandler{
		fileRepo,
		converter,
	}

	// cli get args.
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Please enter a yml file path")
		os.Exit(1)
	}
	pathFile := args[1]

	// query
	query := app.GetParameterDbQuery{pathFile}
	p, error := handler.Handle(query)
	if error != nil {
		log.Fatal("error on yml to parameters")
		os.Exit(1)
	}

	// render
	fmt.Printf("dbs %v", p.Parameters)
}
