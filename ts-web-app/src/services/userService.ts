import { User } from '../models/user';
import { GetCommand, GetCommandOutput, PutCommand } from "@aws-sdk/lib-dynamodb";
import db from '../db';

export const getUserDataById = async(id: string): Promise<GetCommandOutput> => {
    const params = {
        TableName: 'Users',
        Key: {
            ID: id,
        },
    };
    
    const command = new GetCommand(params);
    const result = await db.send(command);

    return result;
}

export const createUser = async (user: User): Promise<User> => {
    const params = {
        TableName: 'Users',
        Item: {
            ID: user.id,
            Name: user.name,
        },
    };

    const command = new PutCommand(params);
    await db.send(command);
    return user;
};