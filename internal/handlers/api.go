package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/Kei-K23/go-dummy-bank-api/api"
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
		GetUserHandler(w, r, db)
	})
	
	mux.HandleFunc("POST /users", func(w http.ResponseWriter, r *http.Request) {
		CreateUserHandler(w, r, db)
	})
	
	mux.HandleFunc("POST /login", func(w http.ResponseWriter, r *http.Request) {
		LoginHandler(w, r, db)
	} )
}