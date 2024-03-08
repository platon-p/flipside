import { ApiService, TokenPairResponse } from "./ApiService";

export const AuthService = {
    isAuth(): boolean {
        return localStorage.getItem('accessToken') !== null;
    },

    logout(): void {
        localStorage.removeItem('accessToken');
        localStorage.removeItem('refreshToken');
    },

    async loginByEmail(email: string, password: string): Promise<string | null> {
        try {
            const tokenPair = await ApiService.Auth.loginByEmail(email, password)
            this._applyTokenPair(tokenPair);
            return null;
        } catch (e: any) {
            return e;
        }
    },

    async register(data: {
        name: string
        nickname: string
        email: string
        password: string
    }): Promise<string | null> {
        try {
            const tokenPair = await ApiService.Auth.register(data);
            this._applyTokenPair(tokenPair);
            return null;
        } catch (e: any) {
            return e;
        }
    },

    _applyTokenPair(tokenPair: TokenPairResponse): void {
        localStorage.setItem('accessToken', tokenPair.access_token);
        localStorage.setItem('refreshToken', tokenPair.refresh_token);
    }
}