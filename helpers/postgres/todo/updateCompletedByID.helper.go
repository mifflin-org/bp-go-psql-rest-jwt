package postgres

import (
	"context"
	"github.com/zerefwayne/go-postgres-rest-docker-boilerplate/config"
	"github.com/zerefwayne/go-postgres-rest-docker-boilerplate/models"
	"log"
)

// UpdateCompletedByID Marks a TO DO Completed
func UpdateCompletedByID(id int) (models.Todo, error) {

	log.Printf("helper	|	marking todo complete by id\n")

	query := "UPDATE todos SET completed=TRUE WHERE id=$1 RETURNING id, content, created_at, completed;"

	var todo models.Todo

	rows := config.DB.QueryRow(context.Background(), query, id)

	if err := rows.Scan(&todo.ID, &todo.Content, &todo.CreatedAt, &todo.Completed); err != nil {
		return todo, err
	}

	return todo, nil

}