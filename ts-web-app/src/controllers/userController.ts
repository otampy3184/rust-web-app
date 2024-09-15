import { Request, Response } from "express";
import { getUserDataById, createUser } from "../services/userService";

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