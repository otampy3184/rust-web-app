// src/handlers/user_handler.rs

use actix_web::{web, HttpResponse};
use serde_json::json;

use crate::services::user_service;
use crate::models::user::User;

// ユーザーを取得するハンドラー
pub async fn get_user(id: web::Path<String>) -> HttpResponse {
    // パスからユーザーIDを取得
    let user_id = id.into_inner();
    
    // サービス層でユーザーを取得
    match user_service::get_user(&user_id).await {
        Ok(user) => HttpResponse::Ok().json(user), // 成功した場合、ユーザー情報をJSONとして返す
        Err(_) => HttpResponse::NotFound().json(json!({"error": "User not found"})), // ユーザーが見つからない場合、404エラーを返す
    }
}

// ユーザーを作成するハンドラー
pub async fn create_user(user: web::Json<User>) -> HttpResponse {
    let user_data = user.into_inner();
    
    // サービス層でユーザーを作成
    match user_service::create_user(user_data).await {
        Ok(user) => HttpResponse::Created().json(user), // 成功した場合、作成したユーザー情報をJSONとして返す
        Err(_) => HttpResponse::InternalServerError().json(json!({"error": "Failed to create user"})), // 失敗した場合、500エラーを返す
    }
}

// ユーザーを更新するハンドラー
pub async fn update_user(id: web::Path<String>, user: web::Json<User>) -> HttpResponse {
    let user_id = id.into_inner();
    let new_data = user.into_inner();
    
    // サービス層でユーザーを更新
    match user_service::update_user(&user_id, new_data.name).await {
        Ok(user) => HttpResponse::Ok().json(user), // 成功した場合、更新後のユーザー情報をJSONとして返す
        Err(_) => HttpResponse::InternalServerError().json(json!({"error": "Failed to update user"})), // 失敗した場合、500エラーを返す
    }
}
