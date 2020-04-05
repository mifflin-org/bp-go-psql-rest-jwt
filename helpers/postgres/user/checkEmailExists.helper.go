package postgres

import (
	"github.com/zerefwayne/go-psql-rest-jwt-docker-boilerplate/config"
	"context"
	"log"
)

// DoesEmailExist Checks if email exists in database
func DoesEmailExist(email string) (bool, error) {

	log.Printf("helper	|	checking email %s exists\n", email)

	query := "SELECT * FROM users WHERE email=$1;"

	if rows, err := config.DB.Query(context.Background(), query, email); err != nil {
		return false, err
	} else {

		// Don't forget to close the rows after calling Query()
		defer rows.Close()

		for rows.Next() {
			// If enters the loop it means email exists in the database, hence return true
			return true, nil
		}

		// If control reaches here, email doesn't exist
		return false, nil
	}

}



