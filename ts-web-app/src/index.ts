import express, { Request, Response } from 'express';

const app = express();
const port = 8080;

interface User {
  id: number;
  name: string;
}

app.get('/', (req: Request, res: Response) => {
  res.send('Hello, TypeScript!');
});

app.get('/user', (req: Request, res: Response) => {
  const user: User = {
    id: 1,
    name: 'Alice',
  };
  res.json(user);
});

app.listen(port, () => {
  console.log(`Server is running on http://localhost:${port}`);
});