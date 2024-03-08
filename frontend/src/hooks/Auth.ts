import { createContext, useContext } from "react";

export interface AuthData {
    isAuth: boolean,
    userId: number | null,
    logout(): void
}

// @ts-ignore
export const authContext = createContext<AuthData>();
export const useAuth = () => useContext(authContext);