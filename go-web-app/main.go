package main

import (
	"log"
	"net/http"
	"go-web-app/server"
)

func main() {
	srv := server.NewServer()
	log.Println("Starting server on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", srv.Router))
}