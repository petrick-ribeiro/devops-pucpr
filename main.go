package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/petrick-ribeiro/devops-pucpr/api/v1"
	"github.com/petrick-ribeiro/devops-pucpr/storage"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Get the environment variables from .env
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	apiPort := os.Getenv("LISTEN_ADDRESS")

	// Start the database connection
	db, err := storage.NewPostgresStorage(
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbName)
	if err != nil {
		log.Fatal(err)
	}

	// Initiate database schemas
	if err := db.Init(); err != nil {
		log.Fatal(err)
	}

	// Define Server settings
	srv := api.NewAPIServer(apiPort, db)

	// Start the server
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
