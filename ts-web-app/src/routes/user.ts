import { Router } from "express";
import { getUserHandler, createUserHandler } from "../controllers/userController";

const router = Router();

router.get('/user', getUserHandler);
router.post('/user', createUserHandler);

export default router;