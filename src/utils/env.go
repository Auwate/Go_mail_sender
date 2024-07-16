package utils

import (
	"errors"
	"os"
)

func EnvSetup() error {

	if os.Getenv("EMAIL") == "" {
		return errors.New("EMAIL environment variable not found")
	}

	if os.Getenv("PASSWORD") == "" {
		return errors.New("PASSWORD environment variable not found")
	}

	return nil

}
