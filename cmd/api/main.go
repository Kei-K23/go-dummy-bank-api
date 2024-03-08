package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Kei-K23/go-dummy-bank-api/internal/handlers"
	"github.com/Kei-K23/go-dummy-bank-api/internal/tools"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Panic("Error loading .env file")
	}

	// Read the PORT from environment variables
	PORT := os.Getenv("PORT")

	// Create a new ServeMux instance
	mux := http.NewServeMux()

	// Connect to the database
	db := tools.ConnectToDB()
	fmt.Println("Successfully connected to database")

	// Register handlers
	handlers.APIHandler(mux, db)

	// Start the server
	fmt.Printf("Server is running on http://localhost%s\n", PORT)
	log.Fatal(http.ListenAndServe(PORT, mux))
}
