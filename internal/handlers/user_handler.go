package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/Kei-K23/go-dummy-bank-api/api"
	"github.com/Kei-K23/go-dummy-bank-api/internal/services"
)


func CreateUserHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
		user := api.User{}

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			api.ErrBadRequest(w, err)
			return
		}

		resUser, err := services.CreateUser(db , &user)

		if err != nil {
			api.ErrInternalServer(w, err)
            return
		}

		resJson, err := json.Marshal(resUser)

		if err != nil {
			api.ErrInternalServer(w, err)
            return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		_ , err = w.Write(resJson)
		if err != nil { 
			api.ErrInternalServer(w , err)
            return
		}
}

func GetUserHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
		id := r.PathValue("id")
		authorization := r.Header.Get("Authorization")

		if authorization == "" {
			newErr := errors.New("authorization header is missing")
			api.ErrUnAuthorizedCredentials(w, newErr)
			return
		}

		authHeader := strings.Split(r.Header.Get("Authorization"), " ")[1] 
		
		if authHeader == "" {
			newErr := errors.New("authorization header is missing")
			api.ErrUnAuthorizedCredentials(w, newErr)
			return
		}

		if id == "" { 
			newErr := errors.New("id parameter is required")
			api.ErrBadRequest(w, newErr)
			return
		}

		user := api.ResForLogin{
			Id : id,
			AccessToken: authHeader,
		}
	
		resUser, err := services.GetUser(db , &user)

		if err != nil {
			api.ErrInternalServer(w, err)
            return
		}

		resJson, err := json.Marshal(resUser)

		if err != nil {
			api.ErrInternalServer(w, err)
            return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		_ , err = w.Write(resJson)
		if err != nil { 
			api.ErrInternalServer(w , err)
            return
		}
	}