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

// response data for create user
type ResForCreateUser struct { 
	Username string `json:"username"`
	Email string `json:"email"`
}

// error
type Error struct { 
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
	ErrInternalServer = func (w http.ResponseWriter, err error)  {
		writeError(w, err.Error(), 500)
	}
	ErrBadRequest = func (w http.ResponseWriter, err error)  {
		writeError(w, err.Error(), 400)
	}
)
