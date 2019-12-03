package query

import (
	"fmt"
	"github.com/floyoops/poc-go-backup/src/backup/domain/adapter"
	"github.com/floyoops/poc-go-backup/src/backup/domain/model"
	"github.com/floyoops/poc-go-backup/src/backup/domain/repository"
)

type GetParametersDbQueryHandler struct {
	Repository repository.FileRepository
	Adapter    adapter.ConvertToParametersDb
}

func (Qh GetParametersDbQueryHandler) Handle(query GetParameterDbQuery) (model.ParametersDb, error)  {
	bytesYml, error := Qh.Repository.GetContent(query.PathFile)
	if error != nil {
		fmt.Errorf(
			"Error GetParametersDbQueryHandler on get content of file '%v'\n, previous: '%v'",
			query.PathFile,
			error,
		)
	}
	return Qh.Adapter.Convert(bytesYml)
}
