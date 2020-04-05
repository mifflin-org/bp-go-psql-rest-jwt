package main

import (
	"context"
	"github.com/zerefwayne/go-postgres-rest-docker-boilerplate/routes"
	"log"
	"net/http"

	"github.com/rs/cors"
	"github.com/zerefwayne/go-postgres-rest-docker-boilerplate/config"
)


func main() {

	log.Println("server	|	initializing")

	// Database setup
	config.ConnectDB()

	// Close the database connection once the main function is finished
	defer config.DB.Close(context.Background())

	// Calls ping method
	config.PingDB()

	// Creates a new Mux Router
	r := routes.NewRouter()

	// This is used to remove CORS that arise when request comes from the same server's another port
	handler := cors.AllowAll().Handler(r)

	// http ListenAndServe is used to listen to requests on 5000 and redirecting them to handler
	if err := http.ListenAndServe(":5000", handler); err != nil {

		// In case there's an error, server is closed and error is logged.
		log.Fatal(err)
	}

}
