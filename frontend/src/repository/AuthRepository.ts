export interface RegisterRequest {
    name: string
    nickname: string
    email: string
    password: string
}

export interface LoginByEmailRequest {
    email: string
    password: string
}

export interface LoginByTokenRequest {
    refresh_token: string
}

export interface TokenPairResponse {
    access_token: string
    refresh_token: string
    expires_at: Date
}

export interface MessageResponse {
    status_code: number
    message: string
}

export const AuthRepository = {
    async register(req: RegisterRequest): Promise<TokenPairResponse | MessageResponse> {
        const endpoint = 'http://localhost/api/auth/register'
        const response = await fetch(endpoint, {
            method: 'POST',
            body: JSON.stringify(req)
        })
        return await response.json() as (TokenPairResponse | MessageResponse)
    },

    async loginByToken(req: LoginByTokenRequest): Promise<TokenPairResponse | MessageResponse> {
        const endpoint = 'http://localhost/api/auth/login-by-token'
        const response = await fetch(endpoint, {
            method: 'POST',
            body: JSON.stringify(req)
        })
        return await response.json() as (TokenPairResponse | MessageResponse)
    },

    async loginByEmail(req: LoginByEmailRequest): Promise<TokenPairResponse | MessageResponse> {
        const endpoint = 'http://localhost/api/auth/login-by-email'
        const response = await fetch(endpoint, {
            method: 'POST',
            body: JSON.stringify(req)
        })
        return await response.json() as (TokenPairResponse | MessageResponse)
    }
};