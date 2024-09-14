import { User } from '../models/user';

export const getUserData = (): User => {
    return {
        id: 1,
        name: 'Alice',
    };
};