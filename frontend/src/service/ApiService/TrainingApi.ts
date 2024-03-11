import { MessageResponse, config } from "./ApiService"

export interface TrainingSummaryResponse {
    id: number
    card_set_id: number
    training_type: string
    status: string
    count_right: number
    count_wrong: number
}

export interface TrainingTask {
    question: string
    question_type: string
    answers: string[]
}

export interface TaskResult {
    answer: string
    is_correct: boolean
}

export const TrainingApi = {
    async getCardSetTrainings(token: string, slug: string): Promise<Array<TrainingSummaryResponse>> {
        const response = await fetch(`${config.baseUrl}/api/cardset/${slug}/trainings`, {
            method: 'GET',
            headers: {
                'Authorization': `Bearer ${token}`
            }
        })
        if (response.ok) {
            return response.json()
        }
        const error = await response.json() as MessageResponse
        throw error.message
    },

    async createTraining(token: string, slug: string, trainingType: string): Promise<TrainingSummaryResponse> {
        const response = await fetch(`${config.baseUrl}/api/cardset/${slug}/trainings?type=${trainingType}`, {
            method: 'POST',
            headers: {
                'Authorization': `Bearer ${token}`
            },
        })
        if (response.ok) {
            return response.json()
        }
        const error = await response.json() as MessageResponse
        throw error.message
    },

    async getNextTask(token: string, trainingId: number): Promise<TrainingTask> {
        const response = await fetch(`${config.baseUrl}/api/training/${trainingId}/next`, {
            method: 'GET',
            headers: {
                'Authorization': `Bearer ${token}`
            }
        })
        if (response.ok) {
            return response.json()
        }
        const error = await response.json() as MessageResponse
        throw error.message
    },

    async submitAnswer(token: string, trainingId: number, answer: string): Promise<TaskResult> {
        const response = await fetch(`${config.baseUrl}/api/training/${trainingId}/submit?answer=${answer}`, {
            method: 'POST',
            headers: {
                'Authorization': `Bearer ${token}`
            }
        })
        if (response.ok) {
            return response.json()
        }
        const error = await response.json() as MessageResponse
        throw error.message
    }
}