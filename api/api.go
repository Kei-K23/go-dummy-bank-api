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
