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

export interface CardSetResponse {
    title: string
    slug: string
    owner_id: number
}

export interface CardRequest {
    question: string
    answer: string
    position: number
}

export interface CardResponse {
    question: string
    answer: string
    position: number
    card_set_id: number
}

export interface MessageResponse {
    status_code: number
    message: string
}

const config = {
    baseUrl: 'http://79.137.204.109',
    auth: '/api/auth',
    cardSet: '/api/cardset',
    cards: '/api/cards',
}

export const ApiService = {
    Auth: {
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
    },
    CardSet: {
        async createCardSet(token: string, title: string, slug: string): Promise<CardSetResponse> {
            const response = await fetch(`${config.baseUrl}${config.cardSet}`, {
                method: 'POST',
                headers: {
                    'Authorization': `Bearer ${token}`
                },
                body: JSON.stringify({
                    title: title,
                    slug: slug,
                })
            })
            if (response.ok) {
                return response.json()
            }
            const error = await response.json() as MessageResponse
            throw error.message
        },

        async getCardSet(slug: string): Promise<CardSetResponse> {
            const response = await fetch(`${config.baseUrl}${config.cardSet}/${slug}`, {
                method: 'GET'
            })
            if (response.status === 200) {
                return response.json()
            }
            const error = await response.json() as MessageResponse
            throw error.message
        },

        async updateCardSet(token: string, slug: string, newTitle: string, newSlug: string): Promise<CardSetResponse> {
            const response = await fetch(`${config.baseUrl}${config.cardSet}/${slug}`, {
                method: 'PUT',
                headers: {
                    'Authorization': `Bearer ${token}`
                },
                body: JSON.stringify({
                    title: newTitle,
                    slug: newSlug,
                })
            })
            if (response.status === 200) {
                return response.json()
            }
            const error = await response.json() as MessageResponse
            throw error.message
        },

        async deleteCardSet(token: string, slug: string): Promise<boolean> {
            const response = await fetch(`${config.baseUrl}${config.cardSet}/${slug}`, {
                method: 'DELETE',
                headers: {
                    'Authorization': `Bearer ${token}`
                }
            })
            if (response.status === 200) {
                return true
            }
            const error = await response.json() as MessageResponse
            throw error.message
        }
    },
    Card: {
        async createCards(token: string, slug: string, request: Array<CardRequest>): Promise<Array<CardResponse>> {
            const response = await fetch(`${config.baseUrl}${config.cards}/${slug}`, {
                method: 'POST',
                headers: {
                    'Authorization': `Bearer ${token}`
                },
                body: JSON.stringify(request)
            })
            if (response.status === 200) {
                return response.json()
            }
            const error = await response.json() as MessageResponse
            throw error.message
        },

        async getCards(slug: string): Promise<Array<CardResponse>> {
            const response = await fetch(`${config.baseUrl}${config.cards}/${slug}`, {
                method: 'GET'
            })
            if (response.status === 200) {
                return response.json()
            }
            const error = await response.json() as MessageResponse
            throw error.message
        },

        async updateCards(token: string, slug: string, request: Array<CardRequest>): Promise<Array<CardResponse>> {
            const response = await fetch(`${config.baseUrl}${config.cards}/${slug}`, {
                method: 'PUT',
                headers: {
                    'Authorization': `Bearer ${token}`
                },
                body: JSON.stringify(request)
            })
            if (response.status === 200) {
                return response.json()
            }
            const error = await response.json() as MessageResponse
            throw error.message
        },

        async deleteCards(token: string, slug: string, positions: Array<number>): Promise<boolean> {
            const response = await fetch(`${config.baseUrl}${config.cards}/${slug}`, {
                method: 'DELETE',
                headers: {
                    'Authorization': `Bearer ${token}`
                },
                body: JSON.stringify(positions)
            })
            if (response.status === 200) {
                return true
            }
            const error = await response.json() as MessageResponse
            throw error.message
        }
    }
}
