package main

import (
    "encoding/json"
    "net/http"
)

type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

func main() {
    http.HandleFunc("/", indexHandler)
    http.HandleFunc("/user", userHandler)
    http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello, Go!"))
}

func userHandler(w http.ResponseWriter, r *http.Request) {
    user := User{
        ID:   1,
        Name: "Alice",
    }
    json.NewEncoder(w).Encode(user)
}