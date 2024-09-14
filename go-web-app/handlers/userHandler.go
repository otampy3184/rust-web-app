package handlers

import (
	"encoding/json"
	"net/http"
	"go-web-app/services"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	user := services.GetUser() // サービス層からユーザー情報を取得
	json.NewEncoder(w).Encode(user) // ユーザー情報をJSON形式でエンコードし、レスポンスとして返す
}