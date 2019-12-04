package infra_adapter

import (
	"github.com/floyoops/poc-go-backup/src/backup/domain/model"
	"github.com/floyoops/poc-go-backup/src/backup/infra/infra_repository"
	"testing"
)

func TestSuccess(t *testing.T)  {
	repo := infra_repository.FileRepository{}
	filePath := "../../../../data/parameters.yml"
	bytesValue, errorRepo := repo.GetContent(filePath)
	if errorRepo != nil {
		t.Errorf("error get content file '%v'", filePath)
	}
	converter := YmlToParameters{}
	result, errorConvert := converter.Convert(bytesValue)
	if errorConvert != nil {
		t.Errorf("Error on convert the content of '%v'", filePath)
	}
	expected := model.ParametersDb{
		DbHost: "db",
		DbPassword: "root_password",
		DbName: "the_db_name",
		DbPort: "",
		DbUser: "root",
	}
	if result != expected {
		t.Errorf("the result %v different of expected %v", result, expected)
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