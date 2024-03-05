import { createContext, useContext, useState } from "react"
import React from "react"

interface TokenPairResponse {
    access_token: string
    refresh_token: string
}

interface MessageResponse {
    status_code: string
    message: string
}

export const AuthService = {
    getAccessToken(): string | undefined {
        return localStorage.getItem('accessToken') ?? undefined;
    },

    logout() {
        localStorage.removeItem('accessToken');
        localStorage.removeItem('refreshToken');
    },

    isAuth(): boolean {
        return this.getAccessToken() !== undefined;
    },

    getMyId(): number | undefined {
        if (!this.isAuth) return undefined;
        const payload = btoa(this.getAccessToken()!).split('.')[1];
        if (!payload) return undefined;
        return JSON.parse(payload)['id'];
    },

    async loginByEmail(email: string, password: string): Promise<string | null> {
        const response = await fetch('http://localhost/api/auth/login-by-email', {
            method: 'POST',
            body: JSON.stringify({
                email: email,
                password: password,
            })
        });
        const json = await response.json();
        if (response.status == 200) {
            const response: TokenPairResponse = json;
            localStorage.setItem('accessToken', response.access_token);
            localStorage.setItem('refreshToken', response.refresh_token);
            return null;
        }
        return json.message;
    }
}

interface AuthData {
    isAuth: boolean,
    userId: number | null,
    logout(): void
}

// @ts-ignore
const authContext = createContext<AuthData>();
const useProvideAuth = () => {
    const [isAuth, setIsAuth] = useState(false);
    const [userId, setUserId] = useState<number | null>(null);

    function logout() {
        setIsAuth(true);
    }

    const init = () => {
        
    }
    return {
        isAuth,
        userId,
        logout,
    }
}

export function AuthProvider({children}: {children: React.ReactNode}) {
    const auth = useProvideAuth();
    return <authContext.Provider value={auth} children={children} />
}

export const useAuth = () => useContext(authContext);