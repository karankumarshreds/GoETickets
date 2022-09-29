package main

import (
	"fmt"
	"log"
	"path/filepath"
	"github.com/karankumarshreds/GoRedis/cmd/app"
	"github.com/karankumarshreds/GoRedis/internal/common"
)

const LOG_FILE_NAME = "logs.log"

func main() {
	absPath, err := filepath.Abs("./log")
	if err != nil {
		log.Println("Error generating absPath", err)
	} 
	
	logger := common.NewLogger(absPath, LOG_FILE_NAME)
	a := app.NewApp(logger)
	a.Init()
	fmt.Println("Starting app")
}