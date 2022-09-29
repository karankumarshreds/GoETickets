package main

import (
	"io"
	"os"
	"fmt"
	"log"
	"path/filepath"
	"github.com/karankumarshreds/GoRedis/cmd/app"
)

const LOG_FILE_NAME = "logs.log"

func main() {
	absPath, err := filepath.Abs("./log")
	if err != nil {
		log.Println("Error generating absPath", err)
	} 
	
	logFile, err := os.OpenFile(fmt.Sprintf("%v\\%v", absPath, LOG_FILE_NAME), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	
	if err != nil {
		log.Println("Error", err)
	}

	// log.Writer() and logFile both implement the Writer method of the io.Writer interface 
	multiWriter := io.MultiWriter(log.Writer(), logFile)
	
	logger := log.New(multiWriter, "LOG:\t", log.Ldate|log.Ltime|log.Llongfile)
	a := app.NewApp(logger)
	a.Init()
	fmt.Println("Starting app")
}