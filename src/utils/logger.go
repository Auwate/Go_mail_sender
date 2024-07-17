package utils

import (
	"log"
	"os"
	"path/filepath"
)

func LoggerSetup(dirPath string) (*os.File, error) {

	_, err := os.Stat(dirPath)

	if os.IsNotExist(err) {
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}

	file, err := os.OpenFile(filepath.Join(dirPath, "log.txt"), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		return nil, err
	}

	log.SetOutput(file)

	return file, nil

}
