package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/amarantec/picpay/internal/models"
	"github.com/amarantec/picpay/internal/utils"
)

func signupUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		log.Printf("Error: %v", err)
		http.Error(w, "could not parse this users", http.StatusBadRequest)
		return
	}

	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = service.SaveUser(ctxTimeout, newUser)
	if err != nil {
		log.Printf("Error: %v", err)
		http.Error(w, "could not create this user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func loginUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("Error: %v", err)
		http.Error(w, "could not parse this user", http.StatusBadRequest)
		return
	}

	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = service.ValidateUserCredentials(ctxTimeout, user)
	if err != nil {
		log.Printf("Error: %v", err)
		http.Error(w, "credentials invalid", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateToken(user.Id, user.Email)
	if err != nil {
		log.Printf("Error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResp, err := json.MarshalIndent(token, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Print("Login successfull!")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
}
