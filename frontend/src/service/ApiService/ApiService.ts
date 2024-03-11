import { AuthApi } from "./AuthApi"
import { CardApi } from "./CardApi"
import { CardSetApi } from "./CardSetApi"
import { ProfileApi } from "./ProfileApi"
import { TrainingApi } from "./TrainingApi"

export const config = {
    baseUrl: 'http://localhost:80',
    auth: '/api/auth',
    cardSet: '/api/cardset',
    cards: '/api/cards',
}

export interface MessageResponse {
    status_code: number
    message: string
}

export const ApiService = {
    Auth: AuthApi,
    CardSet: CardSetApi,
    Card: CardApi,
    Training: TrainingApi,
    Profile: ProfileApi,
}
