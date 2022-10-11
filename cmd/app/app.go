package app

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis/v9"
	"github.com/gorilla/mux"
	"github.com/karankumarshreds/GoRedis/internal/handlers"
)

var ctx = context.Background()

type App struct {
	logger *log.Logger
	router *mux.Router
	redis  *redis.Client
}

const PORT = ":8000"

func NewApp(logger *log.Logger) *App {
	return &App{logger: logger}
}

func (a *App) Init() {
	a.logger.Println("Initializing app")
	a.initRedisClient()
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

func (a *App) initRedisClient() {
	client := redis.NewClient(
		&redis.Options{
			Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       0, // USE DEFAULT DB FOR NOW
		},
	)
	a.redis = client
	_, err := a.redis.Get(ctx, "test").Result()
	if err != redis.Nil {
		a.logger.Fatal("Redis connection failed", err)
	} else {
		a.logger.Println("Redis connection successful")
	}
}
