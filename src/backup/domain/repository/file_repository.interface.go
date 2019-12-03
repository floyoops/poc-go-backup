package repository

type FileRepository interface {
	GetContent(pathFile string) ([]byte, error)
}
