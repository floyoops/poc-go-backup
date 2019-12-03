package adapter

import "github.com/floyoops/poc-go-backup/src/backup/domain/model"

type ConvertToParametersDb interface {
	Convert(bytes []byte) (model.ParametersDb, error)
}
