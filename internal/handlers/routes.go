package handlers

import "net/http"

func SetRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/user-signup", signupUser)
	mux.HandleFunc("/user-login", loginUser)
	mux.HandleFunc("/get-balance/{id}", getBalance)

	return mux
}
