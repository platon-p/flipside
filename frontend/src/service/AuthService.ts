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
        if (response.status == 200){
            const response: TokenPairResponse = json;
            localStorage.setItem('accessToken', response.access_token);
            localStorage.setItem('refreshToken', response.refresh_token);
            return null;
        }
        return json.message;
    }
}