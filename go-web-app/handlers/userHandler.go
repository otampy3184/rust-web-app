package handlers

import (
    "encoding/json"
    "net/http"
    "go-web-app/services"
	"go-web-app/models"
    "github.com/gorilla/mux"
)

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    userId := vars["id"]

    user, err := services.GetUser(userId)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(user)
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
    var user models.User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if user.ID == "" || user.Name == "" {
        http.Error(w, "Invalid input: ID and Name are required", http.StatusBadRequest)
        return
    }

    err = services.CreateUser(&user)
    if err != nil {
        http.Error(w, "Failed to create user: "+err.Error(), http.StatusInternalServerError)
        return
    }

    // レスポンスヘッダーとステータスコードの設定
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    userId := vars["id"]

    var requestData struct {
        Name string `json:"name"`
    }

    err := json.NewDecoder(r.Body).Decode(&requestData)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }
    if requestData.Name == "" {
        http.Error(w, "Name is required", http.StatusBadRequest)
        return
    }

    updatedUser, err := services.UpdateUser(userId, requestData.Name)
    if err != nil {
        http.Error(w, "Failed to update user: "+err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(updatedUser)
}