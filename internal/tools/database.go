package tools

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func ConnectToDB() (*sql.DB) {

	err := godotenv.Load(".env")

	if err!= nil {
        log.Panic("Error loading .env file")
    }
	
	USERNAME := os.Getenv("USERNAME")
	PASSWORD := os.Getenv("PASSWORD")
	
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/go_bank_api", USERNAME, PASSWORD))	

	if err!= nil {
        log.Fatal("Error opening database :" , err)
    }

	createUsersTable(db)

	return db
}

func createUsersTable(db *sql.DB)  {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id VARCHAR(255) PRIMARY KEY NOT NULL,
			username VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL UNIQUE,
			password VARCHAR(255) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)

	if err != nil {
		log.Fatal("Error creating users table :" , err)
	}

	fmt.Println("Users table created")
}