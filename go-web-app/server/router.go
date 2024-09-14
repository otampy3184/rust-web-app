package server

import (
    "github.com/gorilla/mux"
)

// muxはコピーして利用せず、ポインタを参照するのが普通らしい
func NewRouter() *mux.Router { 
    router := mux.NewRouter()
	
    return router
}