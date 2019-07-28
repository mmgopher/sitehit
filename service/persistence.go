package service

import (
	"encoding/gob"
	"os"
)

type GobPersistenceManager struct {
}

func (GobPersistenceManager) Write(filePath string, object interface{}) error {
	file, err := os.Create(filePath)
	if err == nil {
		encoder := gob.NewEncoder(file)
		encoder.Encode(object)
	}
	file.Close()
	return err
}

func (GobPersistenceManager) Read(filePath string, object interface{}) error {
	file, err := os.Open(filePath)

	if err == nil {
		decoder := gob.NewDecoder(file)
		err = decoder.Decode(object)
	}
	file.Close()
	return err
}
