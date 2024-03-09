package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/Kei-K23/go-dummy-bank-api/api"
	"github.com/Kei-K23/go-dummy-bank-api/internal/services"
)

func LoginHandler(db *sql.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		credential := api.ReqForLogin{}

		err := json.NewDecoder(r.Body).Decode(&credential)
		if err != nil {
			api.ErrBadRequest(w, err)
			return
		}

		access , err := services.LoginUser(db , &credential)

		if err != nil {
			api.ErrInternalServer(w, err)
            return
		}

		resJson , err := json.Marshal(access)

		if err != nil {
			api.ErrInternalServer(w , err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(resJson)
		if err != nil { 
			api.ErrInternalServer(w , err)
            return
		}
	}	
}