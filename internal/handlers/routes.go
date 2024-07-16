package handlers

import "net/http"

func SetRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/user-signup", signupUser)
	mux.HandleFunc("/user-login", loginUser)

	return mux
}
