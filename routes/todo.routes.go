package routes

import (
	"github.com/gorilla/mux"
	"github.com/zerefwayne/go-psql-rest-jwt-docker-boilerplate/handlers/todo"
	middleware "github.com/zerefwayne/go-psql-rest-jwt-docker-boilerplate/middleware/auth"
)

func LoadToDoRoutes(r *mux.Router) {

	r.HandleFunc("/api/todos", todo.FetchToDoAll).Methods("GET")

	todoRouter := r.PathPrefix("/api/todo").Subrouter()

	todoRouter.Use(middleware.JWTMiddleware)

	todoRouter.HandleFunc("/", todo.InsertToDoHandler).Methods("POST")
	todoRouter.HandleFunc("/{id}", todo.FetchToDoByID).Methods("GET")
	todoRouter.HandleFunc("/{id}", todo.UpdateCompletedToDoHandler).Methods("PUT")
	todoRouter.HandleFunc("/{id}", todo.DeleteToDoByID).Methods("DELETE")

}
