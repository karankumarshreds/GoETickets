package handlers 

import (
	"log"
	"net/http"
)

type AuthHandlers struct {
	logger *log.Logger
} 

func NewAuthHandlers(logger *log.Logger) *AuthHandlers {
	return &AuthHandlers{logger}
}

func (a *AuthHandlers) Signin(w http.ResponseWriter, r *http.Request) {
	a.logger.Println("Signin handler")
}