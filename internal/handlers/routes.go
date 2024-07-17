package handlers

import (
	"net/http"

	"github.com/amarantec/picpay/internal/middleware"
)

func SetRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	
	mux.HandleFunc("/user-signup", signupUser)
	mux.HandleFunc("/user-login", loginUser)
	mux.HandleFunc("/get-balance/{id}", middleware.Authenticate(getBalance))

	return mux
}
