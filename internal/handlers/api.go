package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/Kei-K23/go-dummy-bank-api/api"
	"github.com/Kei-K23/go-dummy-bank-api/internal/middleware"
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

	mux.Handle("GET /users/{id}", middleware.Logger(middleware.CheckAuthHeader(
		http.HandlerFunc(GetUserHandler(db)),
	)))
	
	mux.Handle("GET /users/{id}/balance", middleware.Logger(middleware.CheckAuthHeader(
		http.HandlerFunc(GetBalanceHandler(db)),
	)))

	mux.Handle("POST /users", middleware.Logger(http.HandlerFunc(CreateUserHandler(db))))

	mux.Handle("PUT /users/{id}", middleware.Logger(http.HandlerFunc(UpdateUserHandler(db))))
	
	mux.Handle("DELETE /users/{id}", middleware.Logger(http.HandlerFunc(DeleteUserHandler(db))))
	
	mux.Handle("POST /login", middleware.Logger(http.HandlerFunc(LoginHandler(db))))
}

