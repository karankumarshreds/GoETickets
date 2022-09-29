package app

import (
	"log"
	"net/http"
)

type App struct {
	logger *log.Logger
}

const PORT = ":8000"

func NewApp(logger *log.Logger) *App {
	return &App{logger}
}

func (a *App) Init() {
	a.logger.Println("Initializing app")
	if err := http.ListenAndServe(PORT, nil); err != nil {
		a.logger.Fatal("Cannot initialize the app")
	}
}