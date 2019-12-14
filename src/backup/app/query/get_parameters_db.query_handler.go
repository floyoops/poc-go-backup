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
	bytesYml, err := Qh.Repository.GetContent(query.PathFile)
	if err != nil {
		return model.ParametersDb{}, fmt.Errorf(
			"Error GetParametersDbQueryHandler on get content of file '%v'\n, previous: '%v'",
			query.PathFile,
			err,
		)
	}
	return Qh.Adapter.Convert(bytesYml)
}
