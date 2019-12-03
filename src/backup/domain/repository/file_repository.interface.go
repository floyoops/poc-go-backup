package domain

type FileRepository interface {
	GetContent(pathFile string) ([]byte, error)
}
