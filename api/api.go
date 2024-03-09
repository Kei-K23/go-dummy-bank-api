package api

import (
	"encoding/json"
	"net/http"
)

// request data for create user
type User struct {
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type DBUser struct {
	Id string `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
	AccessToken string `json:"access_token"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type DBAccount struct {
	Id string `json:"id"`
	AccountNumber string `json:"account_number"`
	Balance int `json:"balance"`
	UserId string `json:"user_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type DBTransitions struct {
	Id string `json:"id"`
	Type string `json:"type"`
	Amount int `json:"amount"`
	AccountId string `json:"account_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type Account struct {
	AccountNumber string `json:"account_number"`
	Balance int `json:"balance"`
	UserId string `json:"user_id"`
}

// response data for create user
type ResForCreateUser struct { 
	Username string `json:"username"`
	Email string `json:"email"`
}

type ReqForLogin struct { 
	Email string `json:"email"`
	Password string `json:"password"`
}

type ResForLogin struct { 
	Id string `json:"id"`
    AccessToken string `json:"access_token"`
}

type Balance struct {
	Username string `json:"username"`
	Balance int `json:"balance"`
}

// error
type Error struct { 
	Message string `json:"message"`
	Code int `json:"code"`
}

type Success struct { 
	Message string `json:"message"`
	Code int `json:"code"`
}



// error message
func writeError (w http.ResponseWriter, msg string, code int) {
	res := Error{
		Message: msg,
		Code: code,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(res)
}

var (
	ErrUserNotFound = func(w http.ResponseWriter, err error)  {
	  writeError(w , err.Error(), 404)	
	} 
    ErrUserExists = func (w http.ResponseWriter, err error)   {
		writeError(w, err.Error(), 409)
	}
    ErrInvalidCredentials = func (w http.ResponseWriter, err error)  {
		writeError(w, err.Error(), 401)
	}
    ErrUnAuthorizedCredentials = func (w http.ResponseWriter, err error)  {
		writeError(w, err.Error(), 403)
	}
	ErrInternalServer = func (w http.ResponseWriter, err error)  {
		writeError(w, err.Error(), 500)
	}
	ErrBadRequest = func (w http.ResponseWriter, err error)  {
		writeError(w, err.Error(), 400)
	}
)
