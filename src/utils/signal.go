package utils

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

var SignalCh = make(chan os.Signal, 1)

func HandleSignal(logFileLocation string) {

	signal.Notify(SignalCh, syscall.SIGINT, syscall.SIGTERM)
	var sig os.Signal = <-SignalCh

	log.Printf("Signal %v received. Cleaning up...", sig.String())
	UploadFile(logFileLocation)

}
