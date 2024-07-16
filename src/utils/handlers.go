package utils

import "net/http"

func FileServerSetup(dirPath string) http.Handler {

	var fileServer http.Handler = http.FileServer(http.Dir(dirPath))
	return fileServer

}
