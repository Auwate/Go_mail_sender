package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/Auwate/go_net_tutorial/src/utils"
	"gopkg.in/gomail.v2"
)

var processes = sync.WaitGroup{}
var LogDirPath = "./src/log"
var StaticDirPath = "./src/static"

func uploadHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var err error = r.ParseMultipartForm(10 ^ 6)

	if err != nil {
		log.Printf("Warning: %v\n", err.Error())
		http.Error(w, "There was an error fulfilling request", http.StatusInternalServerError)
		return
	}

	var email string = r.FormValue("email")
	file, fileHeader, err := r.FormFile("html")

	if err != nil {
		log.Printf("Warning: %v\n", err.Error())
		http.Error(w, "There was an error retrieving your file.", http.StatusInternalServerError)
		return
	}

	defer file.Close()

	dirPath := "./src/file/"

	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		os.MkdirAll(dirPath, 0755)
	}

	absPath, err := filepath.Abs("./src/file/" + fileHeader.Filename)

	if err != nil {
		log.Printf("Warning: %v\n", err.Error())
		http.Error(w, "There was an error finding where to place your file.", http.StatusInternalServerError)
		return
	}

	destination, err := os.Create(absPath)

	if err != nil {
		log.Printf("Warning: %v\n", err.Error())
		http.Error(w, "There was an error creating your file.", http.StatusInternalServerError)
		return
	}

	_, err = io.Copy(destination, file)

	if err != nil {
		log.Printf("Warning: %v\n", err.Error())
		http.Error(w, "There was an error saving your file.", http.StatusInternalServerError)
		return
	}

	destination.Close()
	log.Printf("Successfully saved document %v.\n", fileHeader.Filename)

	go sendMail(absPath, email)
	processes.Add(1)

}

func sendMail(filePath string, to string) {

	file, err := os.ReadFile(filePath)

	if err != nil {
		log.Printf("Warning: %v\n", err.Error())
		os.Remove(filePath)
		return
	}

	message := gomail.NewMessage()
	message.SetHeader("From", os.Getenv("EMAIL"))
	message.SetHeader("To", to)
	message.SetHeader("Subject", "You have mail!")
	message.SetBody("text/html", string(file))
	message.Attach(filePath)

	dialer := gomail.NewDialer(
		"smtp.gmail.com",
		587,
		os.Getenv("EMAIL"),
		os.Getenv("PASSWORD"),
	)

	if err := dialer.DialAndSend(message); err != nil {
		log.Fatalf("Failure: %v\n", err.Error())
		os.Remove(filePath)
	} else {
		log.Printf("Success: Sent email to %v\n", to)
		os.Remove(filePath)
	}

	processes.Done()
}

func main() {

	if err := utils.LoggerSetup(LogDirPath); err != nil {
		log.Fatalf("Failed to open log file: %v", err.Error())
	}
	log.Println("Success: Logger setup.")

	log.Println("Trying: Environment variable setup.")
	if err := utils.EnvSetup(nil); err != nil {
		log.Fatalf("Failed to write env fields: %v", err.Error())
	}
	log.Println("Success: Environment variable setup.")

	log.Println("Trying: File server setup.")
	var fileServer http.Handler = utils.FileServerSetup(StaticDirPath)
	log.Println("Success: File server setup.")

	log.Println("Trying: Handler configurations.")
	http.Handle("/", fileServer)
	http.HandleFunc("/upload", uploadHandler)
	log.Println("Success: Handler configurations.")

	log.Println("Starting server...")
	http.ListenAndServe("localhost:8080", nil)
	log.Println("Cleaning up...")
	processes.Wait()
	log.Println("Finished.")

}
