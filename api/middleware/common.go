package middleware

import (
	"github.com/codegangsta/negroni"
	"net/http"
)

// Common is base common middleware
func Common() negroni.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request,  next http.HandlerFunc) {
		w.Header().Add("Content-Type", "application/json")
		next(w, r)
	}
}