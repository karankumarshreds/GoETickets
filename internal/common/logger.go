package common

import (
	"io"
	"os"
	"log"
)

func NewLogger(absPath string, fileName string) *log.Logger {
	logFile, err := os.OpenFile(absPath+"\\"+fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("ERROR: Cannot open the log file", err)
	}

	// log.Writer() and logFile both implement the Writer method of the io.Writer interface 
	multiWriter := io.MultiWriter(log.Writer(), logFile)
	
	logger := log.New(multiWriter, "LOG:\t", log.Ldate|log.Ltime|log.Llongfile)
	return logger
}