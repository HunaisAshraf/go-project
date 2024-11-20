package main

import (
	"fmt"
	"log"
	"net/http"

	dbconfig "github.com/hunaisashraf/go-auth/internal/api/db/config"
	"github.com/hunaisashraf/go-auth/internal/api/repository"
	"github.com/hunaisashraf/go-auth/internal/api/routes"
	services "github.com/hunaisashraf/go-auth/internal/api/service"
)

// func init() {
// 	err := godotenv.Load(".env")
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// 	slog.Info("env loaded successfully")
// }

func main() {

	conn := dbconfig.DbConfig()

	defer conn.Close()

	repo := repository.NewRepository(conn)
	service := services.NewService(repo)

	r := routes.Router(service)

	fmt.Println("server started localhost:3002")

	err := http.ListenAndServe(":3002", r)
	if err != nil {
		log.Fatal(err)
	}
}
