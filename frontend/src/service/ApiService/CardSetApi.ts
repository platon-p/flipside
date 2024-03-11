import { MessageResponse, config } from "./ApiService";

export interface CardSetResponse {
    title: string
    slug: string
    owner_id: number
}

export const CardSetApi = {
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
};