package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func reverseProxy(target string) http.Handler {
	url, _ := url.Parse(target)
	fmt.Println("sdlflkdsf")

	return httputil.NewSingleHostReverseProxy(url)
}

func main() {
	service1 := "http://auth:3001"
	service2 := "http://posts:3002"

	http.Handle("/api/auth/", http.StripPrefix("/api/auth", reverseProxy(service1)))
	http.Handle("/api/posts/", http.StripPrefix("/api/posts", reverseProxy(service2)))

	fmt.Println("server started in 3000")

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
