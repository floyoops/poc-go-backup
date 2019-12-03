package infra_adapter

import (
	"fmt"
	"github.com/floyoops/poc-go-backup/src/backup/domain/model"
	"gopkg.in/yaml.v2"
)

type YmlToParameters struct {
}

func (c YmlToParameters) Convert(byteValue []byte) (model.ParametersDb, error)  {
	// Unmarshal
	p := model.ParametersDb{}
	err := yaml.Unmarshal([]byte(byteValue), &p)
	if err != nil {
		fmt.Errorf("error: %v", err)
	}

	return p, err
}
