use actix_web::{web, Responder};

// ルート`/`に対応するハンドラ関数
pub async fn index() -> impl Responder {
    "Hello, Rust!" // レスポンスとして文字列を返す
}


// `App`に`/`へのルートを追加するための関数
pub fn init(cfg: &mut web::ServiceConfig) {
    cfg.service(
        // ルート`/`にGETリクエストをマッピング
        web::resource("/")
            .route(web::get().to(index)) // `GET /`リクエストを`index`関数にルーティング
    );
}