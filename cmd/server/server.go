package main

import (
	"fmt"
	"go-project/config"
	userRepository "go-project/internal/api/repository"
	"go-project/internal/api/routes"
	userServices "go-project/internal/api/services"
	jwtauth "go-project/util/auth"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	slog.Info("env loaded successfully")
}

//var Client *mongo.Client
//var UserCollection *mongo.Collection

const SERVER_PORT = "localhost:3000"

func main() {

	client := config.ConnectDB()

	repository := userRepository.NewMongoUserRepository(client, os.Getenv("MONGO_DB_NAME"), os.Getenv("MONGO_COLLECTION_NAME"))

	auth := jwtauth.NewJWTAuth(os.Getenv("ACCESS_TOKEN_SECRET"), os.Getenv("REFRESH_TOKEN_SECRET"))
	userService := userServices.NewUserService(repository, auth)

	r := routes.Router(userService)

	fmt.Println("server running in", SERVER_PORT)
	err := http.ListenAndServe(SERVER_PORT, r)
	if err != nil {
		panic(err)
	}
}
