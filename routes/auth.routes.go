package routes

import (
	"github.com/gorilla/mux"
	"github.com/zerefwayne/go-psql-rest-jwt-docker-boilerplate/handlers/user"
)

func LoadAuthRoutes(r *mux.Router) {
	r.HandleFunc("/api/auth/login", handlers.SignInHandler).Methods("POST")
	r.HandleFunc("/api/auth/signup", handlers.SignUpHandler).Methods("POST")
}