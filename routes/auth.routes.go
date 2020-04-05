package routes

import (
	"../handlers/user"
	"github.com/gorilla/mux"
)

func LoadAuthRoutes(r *mux.Router) {
	r.HandleFunc("/api/auth/login", handlers.SignInHandler).Methods("POST")
	r.HandleFunc("/api/auth/signup", handlers.SignUpHandler).Methods("POST")
}