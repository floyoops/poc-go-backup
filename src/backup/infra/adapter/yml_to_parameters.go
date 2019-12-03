package adapter

import (
	domain "backupYmlToFtp/src/backup/domain/model"
	"fmt"
	"gopkg.in/yaml.v2"
)

type YmlToParameters struct {
}

func (c YmlToParameters) Convert(byteValue []byte) (domain.ParametersDb, error)  {
	// Unmarshal
	p := domain.ParametersDb{}
	err := yaml.Unmarshal([]byte(byteValue), &p)
	if err != nil {
		fmt.Errorf("error: %v", err)
	}

	return p, err
}
