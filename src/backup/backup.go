package main

import (
	"fmt"
	"github.com/floyoops/poc-go-backup/src/backup/app/query"
	"github.com/floyoops/poc-go-backup/src/backup/domain/adapter"
	"github.com/floyoops/poc-go-backup/src/backup/domain/repository"
	"github.com/floyoops/poc-go-backup/src/backup/infra/infra_adapter"
	"github.com/floyoops/poc-go-backup/src/backup/infra/infra_repository"
	"log"
	"os"
)

func main() {
	// bootstrap handler
	var fileRepo repository.FileRepository = infra_repository.FileRepository{}
	var converter adapter.ConvertToParametersDb = infra_adapter.YmlToParameters{}
	var handler = query.GetParametersDbQueryHandler{
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
	q := query.GetParameterDbQuery{pathFile}
	p, error := handler.Handle(q)
	if error != nil {
		log.Fatal("error on yml to parameters")
		os.Exit(1)
	}

	// render
	fmt.Printf("dbs %v", p.Parameters)
}
