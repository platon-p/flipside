import { MessageResponse, config } from "./ApiService"

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

export const CardApi = {
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
};