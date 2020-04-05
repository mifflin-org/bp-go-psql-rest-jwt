package postgres

import (
	"github.com/zerefwayne/go-psql-rest-jwt-docker-boilerplate/config"
	"github.com/zerefwayne/go-psql-rest-jwt-docker-boilerplate/models"
	"context"
	"errors"
	"log"
)

// FetchUserByEmail Fetches user by email
func FetchUserByEmail(email string) (models.User, bool, error) {

	log.Printf("helper	|	fetching user by email\n")

	query := "SELECT * FROM users WHERE email=$1;"

	var user models.User

	if rows, err := config.DB.Query(context.Background(), query, email); err != nil {
		return user, false, err
	} else {

		// Don't forget to close the rows after calling Query()
		defer rows.Close()

		for rows.Next() {

			// Row is obtained in the order which they were created in database
			if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt); err != nil {
				log.Println(err)
				return user, false, err
			}

			return user, true, nil
		}

		return user, false, errors.New("no records found")
	}

}


