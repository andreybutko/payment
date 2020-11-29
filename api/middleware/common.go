package middleware

import (
	"os"
	"github.com/codegangsta/negroni"
	"net/http"
)

// Common is base common middleware
func Common() negroni.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request,  next http.HandlerFunc) {
		name, _ := os.Hostname()
		w.Header().Add("X-Server-Name", name)
		w.Header().Add("Content-Type", "application/json")
		next(w, r)
	}
}