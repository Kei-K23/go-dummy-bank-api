package services

import (
	"database/sql"
	"errors"

	"github.com/Kei-K23/go-dummy-bank-api/api"
	"github.com/google/uuid"
)

func CreateTransition(db *sql.DB, balance int, tType, accountId string) (error) {
	// Perform some basic validation
	if accountId == "" {
		return errors.New("accountId is required")
	}
	if balance == 0 {
		return errors.New("balance is required")
	}
	if tType == "" { 
		return errors.New("type for transition is required")
	}

	id := uuid.New()

	stmt, err := db.Prepare("INSERT INTO transitions (id, type, amount, account_id) VALUES (?, ?, ?, ?)")
	if err != nil {
		return  err 
	}
	defer stmt.Close()

	_, err = stmt.Exec(id, tType, balance, accountId)
	if err != nil {
		return  err 
	}

	return nil
}

func GetTransitions(db *sql.DB, accountId string) ([]api.DBTransitions, error) {
	// Perform some basic validation
	if accountId == "" {
		return nil, errors.New("accountId is required")
	}

	stmt, err := db.Prepare("SELECT * FROM transitions WHERE account_id = ? ORDER BY created_at DESC")
	if err != nil {
		return nil, err 
	}
	defer stmt.Close()

	rows, err := stmt.Query(accountId)
	if err != nil {
		return nil, err 
	}
	defer rows.Close()

	transitions := make([]api.DBTransitions, 0)

	for rows.Next() { 
		var t api.DBTransitions
		err := rows.Scan(&t.Id, &t.Type, &t.Amount, &t.AccountId, &t.CreatedAt, &t.UpdatedAt)
		if err != nil { 
			return  nil, err
		}
		transitions = append(transitions, t)
	}

	if err = rows.Err() ; err != nil { 
		return nil, err
	}

	return transitions, nil
}
