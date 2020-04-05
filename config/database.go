package config

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
)

var (
	// DB Global database
	DB *pgx.Conn
)

func getDatabaseURL() string {

	// Loads the .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	// os.Getenv reads the environment variables and assigns them to the variables
	host := os.Getenv("PG_HOST")
	port := os.Getenv("PG_PORT")
	user := os.Getenv("PG_USER")
	password := os.Getenv("PG_PASSWORD")
	dbname := os.Getenv("PG_DB")

	// TODO Add error checking incase environment variables are missing

	//dbURL is the string constructed from the above variables
	dbURL := "postgres://" + user + ":" + password + "@" + host + ":" + port + "/" + dbname

	return dbURL

}

// ConnectDB : Connects to Postgres Database
func ConnectDB() {

	// Calls the getDatabaseURL function which reads environment variables and constructs a postgres connection URL
	dbURL := getDatabaseURL()

	log.Println("database	|	connecting to database at:", dbURL)

	// Calls pgx.Connect method with context and the database URL previously generated
	if db, err := pgx.Connect(context.Background(), dbURL); err != nil {
		log.Fatal(err)
	} else {
		// Assigns the database connection to the globally declared DB and logs success
		DB = db
		log.Println("database	|	connected successfully")
	}

}

// PingDB Tries to ping the database
func PingDB() {

	if err := DB.Ping(context.Background()); err != nil {
		log.Fatal(err)
	} else {
		log.Println("database	|	ping successful")
	}

}
