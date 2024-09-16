package services

import (
    "context"
    "log"
    "errors"
    "go-web-app/models"

    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/dynamodb"
    "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
    "github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
)

var dbClient *dynamodb.Client

// dynamodbを初期化する
func InitDB() error {
    cfg, err := config.LoadDefaultConfig(context.TODO(),
        config.WithEndpointResolver(aws.EndpointResolverFunc(
            func(service, region string) (aws.Endpoint, error) {
                return aws.Endpoint{
                    URL: "http://localhost:8000", // ローカルDynamoDBのエンドポイント
                }, nil
            })),
        config.WithRegion("ap-northeast-2"),
    )
    if err != nil {
        log.Fatalf("unable to load SDK config, %v", err)
        return err
    }

    dbClient = dynamodb.NewFromConfig(cfg)
    return nil
}

// GETメソッドでユーザーを取得する
func GetUserById(id string) (*models.User, error) {
    input := &dynamodb.GetItemInput{
        TableName: aws.String("Users"),
        Key: map[string]types.AttributeValue{
            "ID": &types.AttributeValueMemberS{Value: id},
        },
    }

    result, err := dbClient.GetItem(context.TODO(), input)
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
    _, err = dbClient.PutItem(context.TODO(), input)
    if err != nil {
        log.Printf("Failed to put item: %v", err)
        return err
    }

    return nil
}