package postgres

import (
	"../../../config"
	"../../../models"
	"context"
	"log"
)

// FetchAll Fetches all todos and returns them
func FetchAll() ([]models.Todo, error) {

	log.Printf("helper	|	fetching all todos\n")

	// Intializing empty Todo array
	var todos []models.Todo

	query := "SELECT * FROM todos;"

	if rows, err := config.DB.Query(context.Background(), query); err != nil {
		return todos, err
	} else {

		// Don't forget to close the rows after calling Query()
		defer rows.Close()

		for rows.Next() {

			var todo models.Todo

			// Row is obtained in the order which they were created in database
			if err := rows.Scan(&todo.ID, &todo.Content, &todo.CreatedAt, &todo.Completed); err != nil {
				log.Println(err)
				return todos, err
			}

			todos = append(todos, todo)
		}

		return todos, nil
	}

}