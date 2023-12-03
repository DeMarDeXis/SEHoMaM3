package savedata

import (
	"os"
)

type SaveFile struct {
	NameFile string
	file     *os.File
}

func NewSaveFile(fileway string) (*SaveFile, error) { // изменено
	file, err := os.Open(fileway)
	if err != nil {
		return nil, err
	}

	return &SaveFile{NameFile: fileway, file: file}, nil
}
