package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/amarantec/picpay/internal/database"
	"github.com/amarantec/picpay/internal/handlers"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../../env/.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	ctx := context.Background()

	serverPort := os.Getenv("SERVER_PORT")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Construir a string de conexão
	connectionString := fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`,
		dbHost, dbPort, dbUser, dbPassword, dbName)

	fmt.Printf("Connection string: %s\n", connectionString)

	Conn, err := database.OpenConnection(ctx, connectionString)
	if err != nil {
		panic(err)
	}
	defer Conn.Close()

	handlers.Configure()
	mux := handlers.SetRoutes()

	port := fmt.Sprintf(":%s", serverPort)
	server := &http.Server{
		Addr:    port,
		Handler: mux,
	}

	fmt.Printf("Server listend on: %s\n", server.Addr)
	log.Fatal(server.ListenAndServe())
}
