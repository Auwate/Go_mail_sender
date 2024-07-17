package utils

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var SignalCh = make(chan os.Signal, 1)

func HandleSignal(file *os.File) {

	signal.Notify(SignalCh, syscall.SIGINT, syscall.SIGTERM)
	var sig os.Signal = <-SignalCh

	log.Printf("Signal %v received. Exiting and cleaning up...\n\n", sig.String())
	err := file.Close()
	if err != nil {
		fmt.Printf("Fail: %v\n", err.Error())
	}

	err = UploadFile(file.Name())
	if err != nil {
		fmt.Printf("Fail: %v\n", err.Error())
	}

	os.Exit(0)

}
