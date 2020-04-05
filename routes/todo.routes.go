package routes

import (
	"../handlers/todo"
	"github.com/gorilla/mux"
)

func LoadToDoRoutes(r *mux.Router) {

	r.HandleFunc("/api/todo", todo.InsertToDoHandler).Methods("POST")
	r.HandleFunc("/api/todo/{id}", todo.FetchToDoByID).Methods("GET")
	r.HandleFunc("/api/todo/{id}", todo.UpdateCompletedToDoHandler).Methods("PUT")
	r.HandleFunc("/api/todo/{id}", todo.DeleteToDoByID).Methods("DELETE")

	r.HandleFunc("/api/todos", todo.FetchToDoAll).Methods("GET")

}
