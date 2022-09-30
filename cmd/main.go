package main

import (
	"os"
	"log"
	"path/filepath"
	"github.com/joho/godotenv"
	"github.com/karankumarshreds/GoRedis/cmd/app"
	"github.com/karankumarshreds/GoRedis/internal/common"
)

const LOG_FILE_NAME = "logs.log"

func main() {

	// environment variabls setup
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Cannot load the env file")
	}

	// logger setup 
	absPath, err := filepath.Abs("./log")
	if err != nil {
		log.Println("Error generating absPath", err)
	} 
	logFile := os.Getenv("LOGS_FILE")
	if logFile == "" {
		log.Fatal("Environment variable for LOGS_FILE not set")
	}
	logger := common.NewLogger(absPath, logFile)

	// app setup
	a := app.NewApp(logger)
	a.Init()
	a.Run()
	
}