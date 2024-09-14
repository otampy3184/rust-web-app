import { User } from '../models/user';
import { DynamoDBDocumentClient, PutCommand } from "@aws-sdk/lib-dynamodb";
import { DynamoDBClient } from "@aws-sdk/client-dynamodb";

// DynamoDBクライアントを初期化
const client = new DynamoDBClient({});
const db = DynamoDBDocumentClient.from(client);

export const getUserData = (): User => {
    return {
        id: 1,
        name: 'Alice',
    };
};

export const createUser = async (user: User): Promise<User> => {
    const params = {
        TableName: 'User',
        Item: {
            ID: user.id,
            Name: user.name,
        },
    };

    const command = new PutCommand(params);
    await db.send(command);
    return user;
};