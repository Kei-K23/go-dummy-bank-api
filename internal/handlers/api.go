package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Kei-K23/go-dummy-bank-api/api"
	"github.com/Kei-K23/go-dummy-bank-api/internal/services"
)

func APIHandler(mux *http.ServeMux, db *sql.DB)  {
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		res , err := json.Marshal(map[string]string{
			"message" : "Hello, world",
		})

		if err!= nil {
			api.ErrInternalServer(w , err)
            return
        }

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		_ , err = w.Write(res)
		if err != nil { 
			api.ErrInternalServer(w , err)
            return
		}
	})

	mux.HandleFunc("GET /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		
		if id == "" { 
			newErr := errors.New("id parameter is required")
			api.ErrBadRequest(w, newErr)
			return
		}

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
	})
	
	mux.HandleFunc("POST /users", func(w http.ResponseWriter, r *http.Request) {
		CreateUser(w, r, db)
	})
	
	mux.HandleFunc("POST /login", func(w http.ResponseWriter, r *http.Request) {
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
	})
}