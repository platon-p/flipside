import { MessageResponse, config } from "./ApiService"
import { CardSetResponse } from "./CardSetApi"

export const ProfileApi = {
    async getUserCards(nickname: string): Promise<Array<CardSetResponse>> {
        const response = await fetch(`${config.baseUrl}/api/users/${nickname}/sets`, {
            method: 'GET'
        })
        if (response.status === 200) {
            return response.json()
        }
        const error = await response.json() as MessageResponse
        throw error.message
    }
}