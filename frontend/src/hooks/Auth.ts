import { createContext, useContext } from "react";

export interface RegisterData {
    name: string
    nickname: string
    email: string
    password: string
}

export interface AuthData {
    isAuth: boolean
    userId: number | null
    nickname: string | undefined
    login(email: string, password: string): Promise<string | null>
    register(data: RegisterData): Promise<string | null>
    logout(): void
}

// @ts-expect-error asdasd
export const authContext = createContext<AuthData>();
export const useAuth = () => useContext(authContext);