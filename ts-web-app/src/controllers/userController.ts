import { Request, Response } from "express";
import { getUserDataById, createUser, updateUser } from "../services/userService";

export const getUserHandler = async (req: Request, res: Response) => {
    const userId = req.params.id;

    try {
        const user = await getUserDataById(userId);

        if (user.Item) {
            res.json(user.Item);
        } else {
            res.status(404).json({error: "User not found"});
        }
    } catch (err){
        res.status(500).json(err)
    }
}

export const createUserHandler = async (req: Request, res: Response) => {
    const user = req.body;
    
    try {
        const savedUser = await createUser(user);
        res.json(savedUser);
    } catch (err) {
        res.status(500).json(err);
    }
}

export const updateUserHander = async (req: Request, res: Response): Promise<void> => {
    const { id } = req.params;
    const { name } = req.body;

    if (!id || !name) {
        res.status(400).send("Invalid request: 'id' and 'name' are required.");
        return;
    }

    try {
        await updateUser(id, name);
        res.status(200).send(`User with ID ${id} updated to name ${name}`);
    } catch (error) {
        res.status(500).send(`Failed to update user name: ${error}`);
    }
}