package server

import (
    "go-web-app/handlers"
    "github.com/gorilla/mux" // Routerをシンプルに管理するため、Muxを利用
)

type Server struct {
    Router *mux.Router
}

// Serverもポインタを参照して利用する
func NewServer() *Server {
    router := mux.NewRouter()

    // ルートを設定
    router.HandleFunc("/", handlers.IndexHandler).Methods("GET")
    router.HandleFunc("/user/", handlers.UserHandler).Methods("GET")

	// Serverインスタンスを使うにあたり、参照用のServerのアドレスを返す
    return &Server{
        Router: router,
    }
}