package app

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/karankumarshreds/GoRedis/internal/handlers"
)

type App struct {
	logger *log.Logger
	router *mux.Router
}

const PORT = ":8000"

func NewApp(logger *log.Logger) *App {
	return &App{logger: logger}
}

func (a *App) Init() {
	a.logger.Println("Initializing app")
	a.router = mux.NewRouter()
	a.initRoutes()
}

func (a *App) Run() {
	a.logger.Printf("Running app on %v", PORT)
	if err := http.ListenAndServe(PORT, a.router); err != nil {
		a.logger.Fatal("Cannot initialize the app")
	}
}

func (a *App) initRoutes() {
	authHandlers := handlers.NewAuthHandlers(a.logger)
	a.router.HandleFunc("/api/signin", authHandlers.Signin).Methods("POST")
}
