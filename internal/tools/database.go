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
	createAccountTable(db)
	createTransitionsTable(db)

	return db
}

func createUsersTable(db *sql.DB)  {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id VARCHAR(255) PRIMARY KEY NOT NULL,
			username VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL UNIQUE,
			password VARCHAR(255) NOT NULL,
			access_token VARCHAR(255) NOT NULL UNIQUE,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)

	if err != nil {
		log.Fatal("Error creating users table :" , err)
	}

	fmt.Println("users table created")
}

func createAccountTable(db *sql.DB)  {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS accounts (
			id VARCHAR(255) PRIMARY KEY NOT NULL,
			account_number VARCHAR(255) NOT NULL UNIQUE,
			balance INTEGER NOT NULL,
			user_id VARCHAR(255) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (user_id) REFERENCES users(id)
		)
	`)

	if err != nil {
		log.Fatal("Error creating account table :" , err)
	}

	fmt.Println("accounts table created")
}

func createTransitionsTable(db *sql.DB)  {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS transitions (
			id VARCHAR(255) PRIMARY KEY NOT NULL,
			type VARCHAR(255) NOT NULL,
			amount INTEGER NOT NULL,
			account_id VARCHAR(255) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (account_id) REFERENCES accounts(id)
		)
	`)

	if err != nil {
		log.Fatal("Error creating transitions table :" , err)
	}

	fmt.Println("transitions table created")
}