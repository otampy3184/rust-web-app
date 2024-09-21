use serde::Serialize;

// `User`構造体を定義し、JSONにシリアライズ可能にするため`Serialize`トレイトを実装
#[derive(Serialize)]
pub struct User {
    pub id: String,
    pub name: String,
}