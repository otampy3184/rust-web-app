use crate::models::user::User;
use crate::db::dynamo::{get_item, put_item, update_item};

pub async fn get_user(id: &str) -> Result<User, &'static str> {
    match get_item(id).await {
        Ok(user) => Ok(user),
        Err(_) => Err("User not found"),
    }
}

pub async fn create_user(user: User) -> Result<User, &'static str> {
    if get_item(&user.id).await.is_ok() {
        return Err("User already exists");
    }

    put_item(user).await
}

pub async fn update_user(id: &str, name: String) -> Result<User, &'static str> {
    let mut user = match get_item(id).await {
        Ok(user) => user,
        Err(_) => return Err("User not found"),
    };

    user.name = name;
    update_item(id, user.name.clone()).await
}
