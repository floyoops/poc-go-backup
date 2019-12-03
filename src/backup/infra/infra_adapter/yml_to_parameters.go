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
		return p, err
	}

	var result map[string]map[string]string
	err2 := yaml.Unmarshal([]byte(byteValue), &result)
	if err2 != nil {
		fmt.Errorf("error %v", err2)
		return p, err2
	}
	p.DbHost = result["parameters"]["database_host"]
	p.DbPort = result["parameters"]["database_port"]
	p.DbName = result["parameters"]["database_name"]
	p.DbUser = result["parameters"]["database_user"]
	p.DbPassword = result["parameters"]["database_password"]

	return p, err
}
