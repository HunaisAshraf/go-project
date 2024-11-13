package main

import (
	"fmt"
	"go-project/internal/api/routes"
	"net/http"
)

const SERVER_PORT = "localhost:3000"

func main() {

	r := routes.Router()

	fmt.Println("server runnning in", SERVER_PORT)
	err := http.ListenAndServe(SERVER_PORT, r)
	if err != nil {
		panic(err)
	}
}
