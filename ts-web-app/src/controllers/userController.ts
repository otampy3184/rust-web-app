import { Request, Response } from "express";
import { User } from "../models/user";
import { getUserData } from "../services/userService";

export const getUser = (req: Request, res: Response) => {
    const user: User = getUserData();
    res.json(user);
}