package postgres

import (
	"../../../config"
	"../../../models"
	"context"
	"log"
)

func InsertUser(user *models.User) error {

	log.Printf("helper	|	inserting user %s\n", user.Email)

	// Template Query to Insert Todo
	query := "INSERT INTO users(id, name, email, password, created_at) VALUES($1, $2, $3, $4, $5) RETURNING id;"

	var lastInsertedID string

	// Runs the Query and assigns the lastInsertedID
	if err := config.DB.QueryRow(context.Background(), query, user.ID, user.Name, user.Email, user.Password, user.CreatedAt).Scan(&lastInsertedID); err != nil {
		return err
	}

	// Successfully inserted the model and assigns the lastReturnedID to todo model
	log.Printf("helper	|	successfully inserted user %s\n", lastInsertedID)

	return nil

}
