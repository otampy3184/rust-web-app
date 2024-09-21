// src/main.rs

mod routes;
mod handlers;
mod services;
mod models;
mod db;

use actix_web::{web, App, HttpServer};

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    // HTTPサーバーを起動し、ルートを設定
    HttpServer::new(|| {
        App::new()
            .configure(routes::init) // ルート設定の初期化
    })
    .bind("127.0.0.1:8080")? // サーバーをバインドするアドレスとポート
    .run() // サーバーを実行
    .await // 非同期タスクとして待機
}