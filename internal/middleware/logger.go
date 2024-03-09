package middleware

import (
	"fmt"
	"net/http"
	"time"
)

// Logger returns a new middleware.Logger that logs each request.
func Logger(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()

        next.ServeHTTP(w, r)
        fmt.Printf("[%s] %s %s\n", r.Method, r.RequestURI, time.Since(start))
    })
}