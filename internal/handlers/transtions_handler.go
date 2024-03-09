package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Kei-K23/go-dummy-bank-api/api"
	"github.com/Kei-K23/go-dummy-bank-api/internal/services"
)


func GetTransitionHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        id := r.PathValue("id")
        
        if id == "" { 
            newErr := errors.New("id parameter is required")
            api.ErrBadRequest(w, newErr)
            return
        }
    
        resTrans, err := services.GetTransitions(db, id)

        if err != nil {
            api.ErrInternalServer(w, err)
            return
        }

        resJson, err := json.Marshal(resTrans)

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
