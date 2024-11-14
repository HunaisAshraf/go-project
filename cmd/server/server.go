package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"go-project/config"
	"go-project/internal/api/routes"
	"go.mongodb.org/mongo-driver/mongo"
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

var Client *mongo.Client
var UserCollection *mongo.Collection

const SERVER_PORT = "localhost:3000"

func main() {

	Client = config.ConnectDB()

	r := routes.Router()

	UserCollection = Client.Database(os.Getenv("MONGO_DB_NAME")).Collection("user")
	fmt.Println("server runnning in", SERVER_PORT)
	err := http.ListenAndServe(SERVER_PORT, r)
	if err != nil {
		panic(err)
	}
}
