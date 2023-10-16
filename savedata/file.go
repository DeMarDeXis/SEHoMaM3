package savedata

import (
	"os"
)

type SaveFile struct {
	Filename string
	File     *os.File
}

func NewSaveFile(filename string) *SaveFile {
	return &SaveFile{
		Filename: filename,
		File:     nil,
	}
}

func (sf *SaveFile) Open() error {
	file, err := os.Open(sf.Filename)
	if err != nil {
		return err
	}
	sf.File = file
	return nil
}
