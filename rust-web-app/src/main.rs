mod routes;
mod models;

use actix_web::{App, HttpServer};

#[actix_web::main] // この属性マクロを利用することでmain関数を非同期で動作するWebサーバーとして定義できる
async fn main() -> std::io::Result<()> {
    HttpServer::new(|| {
        App::new() // App::new()で新しいActix Webアプリケーションを作成し、ルートを構成
            .configure(routes::index::init) // `/`に対応
            .configure(routes::user::init) // `/get_user`に対応
    })
    .bind("127.0.0.1:8080")?
    .run() // サーバーを非同期で実行
    .await // 非同期処理を待機
}