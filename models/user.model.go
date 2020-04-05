package models

import "time"

type User struct {
	ID string `json:"id" db:"id" bson:"id"`
	Name string `json:"name" db:"name" bson:"name"`
	Email string `json:"email" db:"email" bson:"email"`
	Password string `json:"password" db:"password" bson:"password"`
	CreatedAt time.Time `json:"created_at" db:"created_at" bson:"created_at"`
}