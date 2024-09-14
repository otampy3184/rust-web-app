import { DynamoDBClient } from "@aws-sdk/client-dynamodb";
import { DynamoDBDocumentClient } from "@aws-sdk/lib-dynamodb";

const client = new DynamoDBClient({
    region: "ap-northeast-2",
    endpoint: `http://localhost:8000`,
});

const db = DynamoDBDocumentClient.from(client);

export default db;