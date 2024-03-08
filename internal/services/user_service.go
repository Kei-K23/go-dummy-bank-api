package services

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Kei-K23/go-dummy-bank-api/api"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser creates a new user.
func CreateUser(db *sql.DB, user *api.User) (*api.ResForCreateUser, error) {
	// Perform some basic validation
	if user == nil {
		return nil, errors.New("user object is nil")
	}
	if user.Username == "" {
		return nil, errors.New("username is required")
	}
	if user.Email == "" {
		return nil, errors.New("email is required")
	}
	if user.Password == "" {
		return nil, errors.New("password is required")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err // Failed to hash password
	}

	// Generate UUID for the user ID
	id := uuid.New()
	accessToken := uuid.New()
	// Create a prepared statement to insert the user into the database
	stmt, err := db.Prepare("INSERT INTO users (id, username, email, password, access_token) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return nil, err // Failed to prepare statement
	}
	defer stmt.Close()

	// Execute the prepared statement with user data
	_, err = stmt.Exec(id, user.Username, user.Email, string(hashedPassword), accessToken)
	if err != nil {
		return nil, err // Failed to insert user into database
	}

	// Return the created user
	createdUser := &api.ResForCreateUser{
		Username: user.Username,
		Email:    user.Email,
	}
	return createdUser, nil
}

func GetUser(db *sql.DB, user *api.ResForLogin) (*api.DBUser, error) {
	if user == nil {
		return nil, errors.New("user object is nil")
	}
	if user.Id == "" {
		return nil, errors.New("id is required")
	}
	if user.AccessToken == "" {
		return nil, errors.New("accessToken is required")
	}

	stmt, err := db.Prepare("SELECT * FROM users WHERE id = ? AND access_token = ?")
	if err != nil {
		return nil, err // Failed to prepare statement
	}
	defer stmt.Close()

	// Execute the SQL statement to retrieve user information
	var fetchedUser api.DBUser

	err = stmt.QueryRow(user.Id, user.AccessToken).Scan(&fetchedUser.Id, &fetchedUser.Username, &fetchedUser.Email, &fetchedUser.Password, &fetchedUser.AccessToken, &fetchedUser.CreatedAt, &fetchedUser.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			// User with the provided email not found
			fmt.Println(err.Error())
			return nil, errors.New("user not found")
		}
		return nil, err // Other error occurred
	}
	
	return &fetchedUser, nil
}