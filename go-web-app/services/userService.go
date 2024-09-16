package services

import (
    "context"
    "log"
    "errors"
    "go-web-app/models"
    "go-web-app/db"

    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/service/dynamodb"
    "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
    "github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
    "github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
)


// GETメソッドでユーザーを取得する
func GetUser(id string) (*models.User, error) {
    input := &dynamodb.GetItemInput{
        TableName: aws.String("Users"),
        Key: map[string]types.AttributeValue{
            "ID": &types.AttributeValueMemberS{Value: id},
        },
    }

    result, err := db.DbClient.GetItem(context.TODO(), input)
    if err != nil {
        return nil, err
    }

    if result.Item == nil {
        return nil, errors.New("user not found")
    }

    var user models.User
    err = attributevalue.UnmarshalMap(result.Item, &user)
    if err != nil {
        return nil, err
    }

    return &user, nil
}

// POSTメソッドでUserを作成する
func CreateUser(user *models.User) error {
    // User構造体をDynamoDBに保存可能なマップに変換
    item, err := attributevalue.MarshalMap(user)
    if err != nil {
        return err
    }

    // PutItemの入力パラメータを設定
    input := &dynamodb.PutItemInput{
        TableName: aws.String("Users"),
        Item:      item,
    }

    // DynamoDBにPutItemリクエストを送信
    _, err = db.DbClient.PutItem(context.TODO(), input)
    if err != nil {
        log.Printf("Failed to put item: %v", err)
        return err
    }

    return nil
}

// PUTメソッドでUserを更新する
func UpdateUser(id string, name string) (*models.User, error) {
    // 更新対象のユーザーを取得
    user, err := GetUser(id)
    if err != nil {
        return nil, err
    }

    // 更新するフィールドを設定
    update := expression.Set(
        expression.Name("Name"),
        expression.Value(name),
    )

    // 更新リクエストのパラメータを作成
    expr, err := expression.NewBuilder().WithUpdate(update).Build()
    if err != nil {
        return nil, err
    }

    input := &dynamodb.UpdateItemInput{
        TableName: aws.String("Users"),
        Key: map[string]types.AttributeValue{
            "ID": &types.AttributeValueMemberS{Value: id},
        },
        ExpressionAttributeNames:  expr.Names(),
        ExpressionAttributeValues: expr.Values(),
        UpdateExpression:          expr.Update(),
        ReturnValues:              types.ReturnValueUpdatedNew,
    }

    // 更新リクエストを送信
    result, err := db.DbClient.UpdateItem(context.TODO(), input)
    if err != nil {
        return nil, err
    }

    // 更新後のユーザーデータを取得して返す
    err = attributevalue.UnmarshalMap(result.Attributes, user)
    if err != nil {
        return nil, err
    }

    return user, nil
}