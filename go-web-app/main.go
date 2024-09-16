package main

import (
    "log"
    "net/http"
    "go-web-app/server"
    "go-web-app/db"
)

func main() {
    err := db.InitDB()
    if err != nil {
        log.Fatalf("Failed to initialize DynamoDB: %v", err)
    }
	
	router := server.NewRouter()
	log.Println("Starting server on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}