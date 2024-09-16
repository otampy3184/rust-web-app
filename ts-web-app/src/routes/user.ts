import { Router } from "express";
import { getUserHandler, createUserHandler, updateUserHandler } from "../controllers/userController";

const router = Router();

router.get('/user/:id', getUserHandler);
router.post('/user', createUserHandler);
router.put('/user/:id', updateUserHandler);

export default router;