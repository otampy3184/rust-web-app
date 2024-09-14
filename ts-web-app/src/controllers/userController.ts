import { Request, Response } from "express";
import { User } from "../models/user";
import { getUserData, createUser } from "../services/userService";

export const getUserHandler = (req: Request, res: Response) => {
    const user: User = getUserData();
    res.json(user);
}

export const createUserHandler = async (req: Request, res: Response) => {
    const user = req.body;
    try {
        const savedUser = await createUser(user);
        res.json(savedUser);
    } catch (err) {
        res.status(500).json({ err: 'User Creation Failed' });
    }
}