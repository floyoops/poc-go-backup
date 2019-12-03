package infra_repository

import (
	"fmt"
	"io/ioutil"
	"os"
)

type FileRepository struct {
}

func (r FileRepository) exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func (r FileRepository) GetContent(pathFile string) ([]byte, error) {
	// Check if file exist.
	if r.exists(pathFile) == false {
		return make([]byte, 0, 0), fmt.Errorf("the file '%v' does not exist\n", pathFile)
	}
	fmt.Printf("Arg path file yml detected: %v\n", pathFile)

	// Open.
	file, err := os.Open(pathFile)
	if err != nil {
		return make([]byte, 0, 0), fmt.Errorf("error on open file: %v\n", err)
	}
	fmt.Printf("open %v success\n", pathFile)

	// Read
	defer file.Close()
	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return make([]byte, 0, 0), fmt.Errorf("error on read %s\n", err)
	}
	fmt.Printf("Read success")

	return byteValue, err
}
