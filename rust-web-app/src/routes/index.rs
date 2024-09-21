// src/routes/index.rs

use actix_web::web;

use crate::handlers::user_handler::{get_user, create_user, update_user};

// ルート設定を初期化
pub fn init(cfg: &mut web::ServiceConfig) {
    cfg.service(web::resource("/user/{id}").route(web::get().to(get_user))); // ユーザーを取得
    cfg.service(web::resource("/user").route(web::post().to(create_user))); // ユーザーを作成
    cfg.service(web::resource("/user/{id}").route(web::put().to(update_user))); // ユーザーを更新
}