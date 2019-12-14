package infra_repository

import "testing"

func TestGetContentSuccess(t *testing.T)  {
	repo := FileRepository{}
	_, err := repo.GetContent("../../../../data/parameters.yml")
	if err != nil {
		t.Errorf("no %v", err)
	}
}

func TestFileNotExist(t *testing.T)  {
	repo := FileRepository{}
	_, err := repo.GetContent("fake")
	if err != nil && err.Error() != "the file 'fake' does not exist\n" {
		t.Errorf("no")
	}
}
