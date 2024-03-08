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
            localStorage.setItem('accessToken', json.access_token);
            localStorage.setItem('refreshToken', json.refresh_token);
            return null;
        }
        return json.message;
    }
}
