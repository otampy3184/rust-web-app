use aws_sdk_dynamodb::Client;
use aws_sdk_dynamodb::model::AttributeValue;
use aws_config::meta::region::RegionProviderChain;
use crate::models::user::User;
use std::collections::HashMap;

// DynamoDBのクライアントを非同期で初期化
pub async fn get_dynamo_client() -> Client {
    // リージョン情報を環境変数などからロード
    let region_provider = RegionProviderChain::default_provider().or_else("us-east-1");
    let config = aws_config::from_env().region(region_provider).load().await;
    
    Client::new(&config) // DynamoDBクライアントを初期化
}

// ユーザーを取得する関数
pub async fn get_item(id: &str) -> Result<User, &'static str> {
    let client = get_dynamo_client().await;

    let mut key = HashMap::new();
    key.insert("ID".to_string(), AttributeValue::S(id.to_string())); // キーを設定

    match client.get_item().table_name("Users").key("ID", AttributeValue::S(id.to_string())).send().await {
        Ok(output) => {
            if let Some(item) = output.item {
                let user = User {
                    id: item.get("ID").unwrap().as_s().unwrap().to_string(),
                    name: item.get("Name").unwrap().as_s().unwrap().to_string(),
                };
                Ok(user) // ユーザーが見つかった場合
            } else {
                Err("User not found") // ユーザーが見つからなかった場合
            }
        }
        Err(_) => Err("Failed to get item"), // リクエストが失敗した場合
    }
}

// ユーザーを作成する関数
pub async fn put_item(user: User) -> Result<User, &'static str> {
    let client = get_dynamo_client().await;

    let mut item = HashMap::new();
    item.insert("ID".to_string(), AttributeValue::S(user.id.clone()));
    item.insert("Name".to_string(), AttributeValue::S(user.name.clone()));

    // PutItemリクエストを実行
    match client.put_item().table_name("Users").set_item(Some(item)).send().await {
        Ok(_) => Ok(user),
        Err(_) => Err("Failed to put item"),
    }
}

// ユーザーを更新する関数
pub async fn update_item(id: &str, name: String) -> Result<User, &'static str> {
    let client = get_dynamo_client().await;

    let mut key = HashMap::new();
    key.insert("ID".to_string(), AttributeValue::S(id.to_string())); // 更新するユーザーのキー

    let update_expression = "SET Name = :name".to_string();
    let mut expression_attribute_values = HashMap::new();
    expression_attribute_values.insert(":name".to_string(), AttributeValue::S(name.clone()));

    // UpdateItemリクエストを実行
    match client.update_item()
        .table_name("Users")
        .key("ID", AttributeValue::S(id.to_string()))
        .update_expression(update_expression)
        .expression_attribute_values(expression_attribute_values) // 修正: 正しい型で値を渡す
        .send().await {
        Ok(_) => Ok(User { id: id.to_string(), name }), // ユーザーの更新に成功した場合
        Err(_) => Err("Failed to update item"), // リクエストが失敗した場合
    }
}
