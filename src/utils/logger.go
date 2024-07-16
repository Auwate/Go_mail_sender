package utils

import (
	"log"
	"os"
	"path/filepath"
)

func LoggerSetup(dirPath string) error {

	_, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		os.MkdirAll(dirPath, 0755)
	} else {
		return err
	}

	file, err := os.OpenFile(filepath.Join(dirPath, "log.txt"), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		return err
	}

	log.SetOutput(file)

	return nil

}
