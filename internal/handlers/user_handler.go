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


func CreateUserHandler(db *sql.DB) http.HandlerFunc{
     return func(w http.ResponseWriter, r *http.Request) { 
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
}

func UpdateUserHandler(db *sql.DB) http.HandlerFunc{
     return func(w http.ResponseWriter, r *http.Request) { 
        id := r.PathValue("id")

        if id == "" {
            newErr := errors.New("id is required")
            api.ErrBadRequest(w, newErr)
			return
        }

        user := api.ResForCreateUser{}

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			api.ErrBadRequest(w, err)
			return
		}

		resUser, err := services.UpdateUser(db , &user, id)

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
}

func DeleteUserHandler(db *sql.DB) http.HandlerFunc{
     return func(w http.ResponseWriter, r *http.Request) { 
        id := r.PathValue("id")

        if id == "" {
            newErr := errors.New("id is required")
            api.ErrBadRequest(w, newErr)
			return
        }

		resUser, err := services.DeleteUser(db, id)

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
}

// GetUserHandler returns an HTTP handler function that retrieves user data from the database
func GetUserHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        id := r.PathValue("id")

        authHeader := strings.Split(r.Header.Get("Authorization"), " ")[1] 
        
        if id == "" { 
            newErr := errors.New("id parameter is required")
            api.ErrBadRequest(w, newErr)
            return
        }

        user := api.ResForLogin{
            Id:          id,
            AccessToken: authHeader,
        }
    
        resUser, err := services.GetUser(db, &user)

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

        _, err = w.Write(resJson)
        if err != nil { 
            api.ErrInternalServer(w, err)
            return
        }
    }
}

func GetBalanceHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        id := r.PathValue("id")

        authHeader := strings.Split(r.Header.Get("Authorization"), " ")[1] 
        
        if id == "" { 
            newErr := errors.New("id parameter is required")
            api.ErrBadRequest(w, newErr)
            return
        }

        user := api.ResForLogin{
            Id:          id,
            AccessToken: authHeader,
        }
    
        resBalance, err := services.GetBalance(db, &user)

        if err != nil {
            api.ErrInternalServer(w, err)
            return
        }

        resJson, err := json.Marshal(resBalance)

        if err != nil {
            api.ErrInternalServer(w, err)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)

        _, err = w.Write(resJson)
        if err != nil { 
            api.ErrInternalServer(w, err)
            return
        }
    }
}
