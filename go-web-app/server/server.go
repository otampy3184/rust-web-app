package server

import (
    "go-web-app/handlers"
    "github.com/gorilla/mux" // Routerをシンプルに管理するため、Muxを利用
)

// Serverもポインタを参照して利用する
func NewRouter() *mux.Router {
    router := mux.NewRouter()

    // ルートを設定
    router.HandleFunc("/", handlers.IndexHandler).Methods("GET")
    router.HandleFunc("/user/{id}", handlers.GetUserHandler).Methods("GET")
    router.HandleFunc("/user", handlers.CreateUserHandler).Methods("POST")

    return router
}