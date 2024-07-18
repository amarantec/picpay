package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/amarantec/picpay/internal/middleware"
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

func getBalance(w http.ResponseWriter, r *http.Request) {
	var user models.User

	idStr := r.URL.Path[len("/get-balance/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	userId := r.Context().Value(middleware.UserIdKey).(int64)

	user.UserId = userId

	balance, err := service.GetTotalBalanceAccount(ctxTimeout, int64(id))
	if err != nil {
		fmt.Printf("error: %v", err)
		http.Error(w, "could not get this info", http.StatusInternalServerError)
		return
	}

	jsonResp, err := json.MarshalIndent(balance, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
}

/*
	func transfer(w http.ResponseWriter, r *http.Request) {
		var user models.User
		idStr := r.URL.Path[len("/transfer/"):]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid sender_id", http.StatusBadRequest)
			return
		}

		rUser, err := strconv.Atoi(r.URL.Query().Get("destinatario_id"))
		if err != nil {
			http.Error(w, "Invalid destinatario_id", http.StatusBadRequest)
			return
		}

		value, err := strconv.ParseFloat(r.URL.Query().Get("valor"), 64)
		if err != nil {
			http.Error(w, "Invalid valor", http.StatusBadRequest)
			return
		}
		ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		userId := r.Context().Value(middleware.UserIdKey).(int64)
		user.UserId = userId

		err = service.Transfer(ctxTimeout, int64(id), int64(rUser), value)
		if err != nil {
			log.Printf("Error: %v", err)
			http.Error(w, "could not transfer this value", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
*/
func transfer(w http.ResponseWriter, r *http.Request) {
	var user models.User
	idStr := r.URL.Path[len("/transfer/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid sender_id", http.StatusBadRequest)
		return
	}

	rUser, err := strconv.Atoi(r.URL.Query().Get("destinatario_id"))
	if err != nil {
		http.Error(w, "Invalid destinatario_id", http.StatusBadRequest)
		return
	}

	value, err := strconv.ParseFloat(r.URL.Query().Get("valor"), 64)
	if err != nil {
		http.Error(w, "Invalid valor", http.StatusBadRequest)
		return
	}
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	userId := r.Context().Value(middleware.UserIdKey).(int64)
	user.UserId = userId

	authorized, err := middleware.CheckExternService()
	if err != nil {
		fmt.Printf("Error: %v", err)
		http.Error(w, "Error when get extern authorization", http.StatusInternalServerError)
		return
	}
	if !authorized {
		fmt.Printf("Error transfer not authorized: %v", err)
		http.Error(w, "Transfer not authorized", http.StatusInternalServerError)
		return
	}

	err = service.Transfer(ctxTimeout, int64(id), int64(rUser), value)
	if err != nil {
		log.Printf("Error: %v", err)
		http.Error(w, "could not transfer this value", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
