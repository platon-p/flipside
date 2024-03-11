import { ApiService, TokenPairResponse } from "./ApiService/ApiService";

export const AuthService = {
    getUserId(): number | null {
        const token = this.getToken()?.split('.')[1];
        if (!token) {
            return null;
        }
        const payload = JSON.parse(atob(token));
        return payload.id;
    },

    getNickname(): string | undefined {
        const token = this.getToken()?.split('.')[1];
        if (!token) {
            return undefined;
        }
        const payload = JSON.parse(atob(token));
        return payload.nickname;
    },

    getToken(): string | null {
        return localStorage.getItem('accessToken');
    },

    isAuth(): boolean {
        return this.getToken() !== null && this.getUserId() !== null;
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

    async loginByRefreshToken(): Promise<string | null> {
        try {
            const refreshToken = localStorage.getItem('refreshToken');
            if (!refreshToken) {
                return 'No refresh token';
            }
            const tokenPair = await ApiService.Auth.loginByToken(refreshToken);
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