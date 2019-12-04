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
	var result map[string]map[string]string
	err := yaml.Unmarshal([]byte(byteValue), &result)
	if err != nil {
		return model.ParametersDb{}, fmt.Errorf("error %v", err)
	}
	// Adapt to model
	p := model.ParametersDb{
		DbHost: result["parameters"]["database_host"],
		DbPort: result["parameters"]["database_port"],
		DbName: result["parameters"]["database_name"],
		DbUser: result["parameters"]["database_user"],
		DbPassword: result["parameters"]["database_password"],
	}
	return p, err
}
