import { Router } from "express";
import { getUserHandler, createUserHandler, updateUserHander } from "../controllers/userController";

const router = Router();

router.get('/user/:id', getUserHandler);
router.post('/user', createUserHandler);
router.put('/user/:id', updateUserHander);

export default router;