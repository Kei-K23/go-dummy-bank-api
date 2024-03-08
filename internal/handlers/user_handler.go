package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/Kei-K23/go-dummy-bank-api/api"
	"github.com/Kei-K23/go-dummy-bank-api/internal/services"
)


func CreateUser(w http.ResponseWriter, r *http.Request, db *sql.DB) {
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