import { Router } from "express";
import { getUserHandler, createUserHandler } from "../controllers/userController";

const router = Router();

router.get('/user/:id', getUserHandler);
router.post('/user', createUserHandler);

export default router;