package utils

import (
	"log"
	"os"
	"path/filepath"
)

func LoggerSetup(dirPath string) error {

	_, err := os.Stat(dirPath)

	if os.IsNotExist(err) {
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	file, err := os.OpenFile(filepath.Join(dirPath, "log.txt"), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	log.Println(file.Name())
	if err != nil {
		return err
	}

	log.Println(os.Getwd())
	os.Chdir("./src")
	log.Println(os.Getwd())
	os.Chdir("./log")
	log.Println(os.Getwd())

	log.SetOutput(file)

	return nil

}
