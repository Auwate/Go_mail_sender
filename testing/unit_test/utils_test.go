package main

import (
	"net/http"
	"os"
	"testing"

	"github.com/Auwate/go_net_tutorial/src/utils"
)

func TestCorrectLoggerSetup(t *testing.T) {

	dirPath := "./log"

	_, err := utils.LoggerSetup(dirPath)

	if err != nil {
		t.Errorf("Failed: %v\n", err.Error())
	}

	if _, err := os.Open(dirPath); err != nil {
		t.Errorf("Failed: %v\n", err.Error())
	}

}

func TestCorrectFileServerSetup(t *testing.T) {

	dirPath := "./testing_files"
	var fs http.Handler = utils.FileServerSetup(dirPath)

	if fs == nil {
		t.Error("Failed: Unknown error\n")
	}

}

func TestEnvSetup(t *testing.T) {

	os.Setenv("EMAIL", "testing@email.com")
	os.Setenv("PASSWORD", "abc123")

	if err := utils.EnvSetup(); err != nil {
		t.Errorf("Failed: %v\n", err.Error())
	}

	if os.Getenv("EMAIL") != "testing@email.com" {
		t.Errorf("Failed: EMAIL should not be %v\n", os.Getenv("EMAIL"))
	}

	if os.Getenv("PASSWORD") != "abc123" {
		t.Errorf("Failed: PASSWORD should not be %v\n", os.Getenv("PASSWORD"))
	}

}
