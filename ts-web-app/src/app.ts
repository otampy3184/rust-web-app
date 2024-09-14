import express from 'express';
import dotenv from 'dotenv';
import indexRoutes from './routes/index';
import userRoutes from './routes/user';

dotenv.config();

const app = express();

app.use(express.json());

app.use('/', indexRoutes);
app.use('/user', userRoutes);

export default app;