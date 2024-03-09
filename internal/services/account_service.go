package services

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/Kei-K23/go-dummy-bank-api/api"
	"github.com/google/uuid"
)

func CreateAccount(db *sql.DB, userId string) (*api.Account, error) {
	// Perform some basic validation
	if userId == "" {
		return nil, errors.New("user id is nil")
	}

	id := uuid.New()
	accountNumber := uuid.New()

	stmt, err := db.Prepare("INSERT INTO accounts (id, account_number, balance, user_id) VALUES (?, ?, ?, ?)")
	if err != nil {
		return nil, err // Failed to prepare statement
	}
	defer stmt.Close()

	_, err = stmt.Exec(id, accountNumber, 500, userId)
	if err != nil {
		return nil, err // Failed to insert account into database
	}

	createdAcc := &api.Account{
		AccountNumber: accountNumber.String(),
		Balance: 500,
		UserId: userId,
	}
	return createdAcc, nil
}


func UpdateAccount(db *sql.DB, balance int, id string) (*api.Success, error) {
	// Perform some basic validation
	if id == "" {
		return nil, errors.New("id is required")
	}
	if balance == 0 {
		return nil, errors.New("balance is required")
	}

	stmt, err := db.Prepare("UPDATE accounts SET balance=? WHERE id=?")
	if err != nil {
		return nil, err 
	}
	defer stmt.Close()

	updatedAcc := &api.Success{
		Message: fmt.Sprintf("Updated balance for account id %s is %d", id, balance),
		Code: http.StatusOK,
	} 
	_, err = stmt.Exec(balance, id)
	if err != nil {
		return nil, err 
	}

	return updatedAcc, nil
}

func DepositAccount(db *sql.DB, balance int, id string) (*api.Success, error) {
	// Perform some basic validation
	if id == "" {
		return nil, errors.New("id is required")
	}
	if balance == 0 {
		return nil, errors.New("balance is required")
	}

	existingBalance, err := getBalance(db , id)

	if err!= nil {
        return nil, err
    }

	stmt, err := db.Prepare("UPDATE accounts SET balance=? WHERE id=?")
	if err != nil {
		return nil, err 
	}
	defer stmt.Close()

	newBalance := existingBalance + balance

	_, err = stmt.Exec(newBalance, id)
	if err != nil {
		return nil, err 
	}

	updatedAcc := &api.Success{
		Message: fmt.Sprintf("Updated balance after deposit for account id %s is %d", id, newBalance),
		Code: http.StatusOK,
	} 
	return updatedAcc, nil
}

func WithdrawAccount(db *sql.DB, balance int, id string) (*api.Success, error) {
	// Perform some basic validation
	if id == "" {
		return nil, errors.New("id is required")
	}
	if balance == 0 {
		return nil, errors.New("balance is required")
	}

	existingBalance, err := getBalance(db , id)

	if err!= nil {
        return nil, err
    }

	if existingBalance < balance {
		return nil, errors.New("insufficient balance")
	}

	stmt, err := db.Prepare("UPDATE accounts SET balance=? WHERE id=?")
	if err != nil {
		return nil, err 
	}
	defer stmt.Close()

	newBalance := existingBalance - balance

	_, err = stmt.Exec(newBalance, id)
	if err != nil {
		return nil, err 
	}

	updatedAcc := &api.Success{
		Message: fmt.Sprintf("Updated balance after deposit for account id %s is %d", id, newBalance),
		Code: http.StatusOK,
	} 
	return updatedAcc, nil
}


func DeleteAccount(db *sql.DB, id string) (*api.Success, error) {

	// Create a prepared statement to update the user
	stmt, err := db.Prepare("DELETE FROM accounts WHERE id=?")
	if err != nil {
		return nil, err // Failed to prepare statement
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return nil, err
	}

	// Return the created user
	res := &api.Success{
		Message: fmt.Sprintf("account with id %s deleted", id),
		Code: http.StatusOK,
	}
	return res, nil
}

func GetAccount(db *sql.DB, obj *api.ResForLogin) (*api.DBAccount, error) {
	if obj == nil {
		return nil, errors.New("obj is nil")
	}
	if obj.Id == "" {
		return nil, errors.New("id is required")
	}
	if obj.AccessToken == "" {
		return nil, errors.New("accessToken is required")
	}

	stmt, err := db.Prepare(`
		SELECT a.*
		FROM accounts a
		INNER JOIN users u ON a.user_id = u.id
		WHERE a.id = ? AND u.access_token = ?
	`)
	if err != nil {
		return nil, err // Failed to prepare statement
	}
	defer stmt.Close()

	var fetchedAcc api.DBAccount

	err = stmt.QueryRow(obj.Id, obj.AccessToken).Scan(&fetchedAcc.Id, &fetchedAcc.AccountNumber, &fetchedAcc.Balance, &fetchedAcc.UserId, &fetchedAcc.CreatedAt, &fetchedAcc.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println(err.Error())
			return nil, errors.New("account not found")
		}
		return nil, err 
	}
	return &fetchedAcc, nil
}

func getBalance(db *sql.DB, id string) (int, error) {
	var balance *int
	stmt, err := db.Prepare("SELECT balance FROM accounts WHERE id = ?")

	if err!= nil {
        return 0, errors.New("error preparing statement")
    }

	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(&balance)

	if err != nil {
		return 0, errors.New("error while executing")
	}

	return *balance, nil
}