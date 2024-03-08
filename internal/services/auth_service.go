package services

import (
	"database/sql"
	"errors"

	"github.com/Kei-K23/go-dummy-bank-api/api"
	"golang.org/x/crypto/bcrypt"
)


func LoginUser(db *sql.DB, user *api.ReqForLogin) (*api.ResForLogin, error) {
	// Check if the user object is nil
	if user == nil {
		return nil, errors.New("user object is nil")
	}

	// Check if email is provided
	if user.Email == "" {
		return nil, errors.New("email is required")
	}

	// Check if password is provided
	if user.Password == "" {
		return nil, errors.New("password is required")
	}

	// Prepare SQL statement to select user by email
	stmt, err := db.Prepare("SELECT * FROM users WHERE email = ?")
	if err != nil {
		return nil, err // Failed to prepare statement
	}
	defer stmt.Close()

	// Execute the SQL statement to retrieve user information
	var fetchedUser api.DBUser
	err = stmt.QueryRow(user.Email).Scan(&fetchedUser.Id, &fetchedUser.Username, &fetchedUser.Email, &fetchedUser.Password, &fetchedUser.AccessToken, &fetchedUser.CreatedAt, &fetchedUser.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			// User with the provided email not found
			return nil, errors.New("user not found")
		}
		return nil, err // Other error occurred
	}

	// Compare the hashed password from the database with the provided password
	if err := bcrypt.CompareHashAndPassword([]byte(fetchedUser.Password), []byte(user.Password)); err != nil {
		// Passwords do not match
		return nil, errors.New("incorrect password")
	}

	// Passwords match, create response object
	response := &api.ResForLogin{
		Id: fetchedUser.Id,
		AccessToken: fetchedUser.AccessToken,
	}

	return response, nil
}