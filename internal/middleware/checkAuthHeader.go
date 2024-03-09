package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/Kei-K23/go-dummy-bank-api/api"
)

// CheckAuthHeader is a middleware function that checks if the request contains an authorization header
// If the authorization header is present, it will be checked for the Bearer type
// If the authorization header is not present or is not of the Bearer type, an error will be returned
// If the authorization header is valid, the request will be forwarded to the next middleware or the endpoint handler
func CheckAuthHeader(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")

        // Check if the authorization header is present
        if authHeader == "" {
            // Return an unauthorized error if the authorization header is not present
            newErr := errors.New("authorization header is missing")
            api.ErrUnAuthorizedCredentials(w, newErr)
            return
        }

        // Split the authorization header into its parts
        parts := strings.Split(authHeader, " ")
        if len(parts)!= 2 || parts[0]!= "Bearer" {
            // Return an unauthorized error if the authorization header is not of the Bearer type
            newErr := errors.New("error when parsing authorization header")
            api.ErrUnAuthorizedCredentials(w, newErr)
            return
        }

        // Forward the request to the next middleware or the endpoint handler if the authorization header is valid
        next.ServeHTTP(w, r)
    })
}