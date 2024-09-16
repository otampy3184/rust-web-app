package db

import (
    "context"
    "log"

    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var DbClient *dynamodb.Client

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

    DbClient = dynamodb.NewFromConfig(cfg)
    return nil
}