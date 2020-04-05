package postgres

import (
	"context"
	"../../../config"
	"log"
)

// DeleteByID Deleted ToDo By ID
func DeleteByID(id int) error {

	log.Printf("helper	|	deleting todo by id\n")

	query := "DELETE FROM todos WHERE id=$1;"

	if _, err := config.DB.Exec(context.Background(), query, id); err != nil {
		return err
	}

	return nil

}

