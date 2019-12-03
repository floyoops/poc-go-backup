package domain

import domain "backupYmlToFtp/src/backup/domain/model"

type ConvertToParametersDb interface {
	Convert(bytes []byte) (domain.ParametersDb, error)
}
