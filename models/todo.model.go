package models

import (
	"time"
)

//Todo Model for todo
type Todo struct {
	ID        int       `db:"id" bson:"id" json:"id"`
	Content   string    `db:"content" bson:"content" json:"content"`
	CreatedAt time.Time `db:"created_at" bson:"created_at" json:"createdAt"`
	Completed bool      `db:"completed" bson:"completed" json:"completed"`
}
