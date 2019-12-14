package main

import (
	"flag"
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
		Repository: fileRepo,
		Adapter:    converter,
	}

	// cli get args.
	var pathFile string
	flag.StringVar(&pathFile,"f", "data/parameters.yml", "help message for ip")
	flag.Parse()
	fmt.Println("pathFile has value ", pathFile)

	// query
	q := query.GetParameterDbQuery{PathFile: pathFile}
	p, err := handler.Handle(q)
	if err != nil {
		log.Fatal("err on yml to parameters, detail: ", err)
		os.Exit(1)
	}

	// render
	fmt.Printf("dbs %v", p)
}
