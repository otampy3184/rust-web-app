package services

import "go-web-app/models" 

func GetUser() models.User {
    // Userモデルのインスタンスを作成して返すだけ
    return models.User{
        ID:   1,
        Name: "Alice",
    }
}