package postgres

import (
	"context"
	"errors"
	"github.com/zerefwayne/go-postgres-rest-docker-boilerplate/config"
	"github.com/zerefwayne/go-postgres-rest-docker-boilerplate/models"
	"log"
)

// FetchByID Fetches todo by ID
func FetchByID(id int) (models.Todo, error) {

	log.Printf("helper	|	fetching todo by id\n")

	query := "SELECT * FROM todos WHERE id=$1;"

	var todo models.Todo

	if rows, err := config.DB.Query(context.Background(), query, id); err != nil {
		return todo, err
	} else {

		// Don't forget to close the rows after calling Query()
		defer rows.Close()

		for rows.Next() {

			// Row is obtained in the order which they were created in database
			if err := rows.Scan(&todo.ID, &todo.Content, &todo.CreatedAt, &todo.Completed); err != nil {
				log.Println(err)
				return todo, err
			}

			return todo, nil
		}

		return todo, errors.New("no records found")
	}

}


