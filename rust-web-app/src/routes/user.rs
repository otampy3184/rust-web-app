use actix_web::{web, Responder, HttpResponse};
use crate::models::user::User;

// ルート`/user`に対応するハンドラ関数
pub async fn get_user() -> impl Responder {
    // `User`構造体のインスタンスを作成
    let user = User {
        id: 1,
        name: "Alice".to_string(),
    };

    // JSON形式で200ステータスのHTTPレスポンスを返す
    HttpResponse::Ok().json(user)
}

// `App`に`/user`ルートを追加するための関数
pub fn init(cfg: &mut web::ServiceConfig) {
    cfg.service(
        web::resource("/user")
            .route(web::get().to(get_user))
    );
}