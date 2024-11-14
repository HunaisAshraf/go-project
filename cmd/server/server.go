package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"go-project/config"
	userRepository "go-project/internal/api/repository"
	"go-project/internal/api/routes"
	userServices "go-project/internal/api/services"
	"log"
	"log/slog"
	"net/http"
	"os"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	slog.Info("env loaded successfully")
}

//var Client *mongo.Client
//var UserCollection *mongo.Collection

const SERVER_PORT = "localhost:3000"

func main() {

	client := config.ConnectDB()

	repository := userRepository.NewMongoUserRepository(client, os.Getenv("MONGO_DB_NAME"), os.Getenv("MONGO_COLLECTION_NAME"))

	userService := userServices.NewUserService(repository)

	r := routes.Router(userService)

	fmt.Println("server running in", SERVER_PORT)
	err := http.ListenAndServe(SERVER_PORT, r)
	if err != nil {
		panic(err)
	}
}
