package infra_adapter

import (
	"fmt"
	"github.com/floyoops/poc-go-backup/src/backup/domain/model"
	"github.com/floyoops/poc-go-backup/src/backup/infra/infra_repository"
	"testing"
)

var Expected  = model.ParametersDb{
	DbHost: "db",
	DbPassword: "root_password",
	DbName: "the_db_name",
	DbPort: "",
	DbUser: "root",
}

func FileYmlToModelExpect(filePath string, expected model.ParametersDb) (bool, error)  {
	repo := infra_repository.FileRepository{}
	bytesValue, errorRepo := repo.GetContent(filePath)
	if errorRepo != nil {
		return false, fmt.Errorf("error get content file '%v'", filePath)
	}
	converter := YmlToParameters{}
	result, errorConvert := converter.Convert(bytesValue)
	if errorConvert != nil {
		return false, fmt.Errorf("error on convert the content of '%v'", filePath)
	}
	if result != expected {
		return false, fmt.Errorf("the result %v different of expected %v", result, Expected)
	}
	return true, nil
}

func TestSuccess(t *testing.T)  {
	_, err := FileYmlToModelExpect("../../../../data/parameters.yml", Expected)
	if err != nil {
		t.Error(err)
	}
}

func TestYmlDisorderSuccess(t *testing.T)  {
	_, err := FileYmlToModelExpect("../../../../data/parameters_disorder.yml", Expected)
	if err != nil {
		t.Error(err)
	}
}

func TestYulMinimalSuccess(t *testing.T)  {
	_, err := FileYmlToModelExpect("../../../../data/parameters_minimal.yml", Expected)
	if err != nil {
		t.Error(err)
	}
}

func TestYmlEmptyFail(t *testing.T)  {
	_, err := FileYmlToModelExpect("../../../../data/parameters_empty.yml", Expected)
	if err == nil {
		t.Error("expected error not found")
	}
}

func TestYmlIncompleteFail(t *testing.T)  {
	_, err := FileYmlToModelExpect("../../../../data/parameters_incomplete.yml", Expected)
	errorExpected := "the result {db    } different of expected {db  the_db_name root root_password}"
	if err != nil && err.Error() != errorExpected {
		t.Errorf("error message '%v' is different of expected message: '%v'", err.Error(), errorExpected)
	}
}

func TestFail(t *testing.T)  {
	byteValues := []byte("ABC€")
	converter := YmlToParameters{}
	_, errorConvert := converter.Convert(byteValues)
	if errorConvert == nil {
		t.Errorf("expected error not found")
	}
}
