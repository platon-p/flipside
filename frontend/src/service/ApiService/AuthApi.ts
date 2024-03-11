import { MessageResponse, config } from "./ApiService"

export interface RegisterRequest {
    name: string
    nickname: string
    email: string
    password: string
}

export interface TokenPairResponse {
    access_token: string
    refresh_token: string
    expires_in: Date
}

export const AuthApi = {
    async register(request: RegisterRequest): Promise<TokenPairResponse> {
        const response = await fetch(`${config.baseUrl}${config.auth}/register`, {
            method: 'POST',
            body: JSON.stringify(request)
        })
        if (response.status === 200) {
            return response.json()
        }
        const error = await response.json() as MessageResponse
        throw error.message
    },

    async loginByEmail(email: string, password: string): Promise<TokenPairResponse> {
        const response = await fetch(`${config.baseUrl}${config.auth}/login-by-email`, {
            method: 'POST',
            body: JSON.stringify({
                email: email,
                password: password,
            })
        })
        if (response.status === 200) {
            return response.json()
        }
        const error = await response.json() as MessageResponse
        throw error.message
    },

    async loginByToken(token: string): Promise<TokenPairResponse> {
        const response = await fetch(`${config.baseUrl}${config.auth}/login-by-token`, {
            method: 'POST',
            body: JSON.stringify({
                refresh_token: token,
            })
        })
        if (response.status === 200) {
            return response.json()
        }
        const error = await response.json() as MessageResponse
        throw error.message
    }
};