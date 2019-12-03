package app

import (
	d2 "backupYmlToFtp/src/backup/domain/adapter"
	d3 "backupYmlToFtp/src/backup/domain/model"
	d1 "backupYmlToFtp/src/backup/domain/repository"
	"fmt"
)

type GetParametersDbQueryHandler struct {
	Repository d1.FileRepository
	Adapter    d2.ConvertToParametersDb
}

func (Qh GetParametersDbQueryHandler) Handle(query GetParameterDbQuery) (d3.ParametersDb, error)  {
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
