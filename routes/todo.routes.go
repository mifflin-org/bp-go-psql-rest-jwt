package routes

import (
	"github.com/gorilla/mux"
	"github.com/zerefwayne/go-postgres-rest-docker-boilerplate/handlers/todo"
)

func LoadToDoRoutes(r *mux.Router) {

	r.HandleFunc("/api/todo", todo.InsertToDoHandler).Methods("POST")
	r.HandleFunc("/api/todo/{id}", todo.FetchToDoByID).Methods("GET")
	r.HandleFunc("/api/todo/{id}", todo.UpdateCompletedToDoHandler).Methods("PUT")
	r.HandleFunc("/api/todo/{id}", todo.DeleteToDoByID).Methods("DELETE")

	r.HandleFunc("/api/todos", todo.FetchToDoAll).Methods("GET")

}
