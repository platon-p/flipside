import { ApiService } from "@/service/ApiService";

export interface TrainingSummary {
    id: number,
    card_set_id: number,
    training_type: string,
    status: string,
    count_right: number,
    count_wrong: number,
}

export const TrainingRepository = {
    async getCardSetTrainings(slug: string): Promise<TrainingSummary[]> {
        const token = localStorage.getItem('accessToken')!;
        const trainings = await ApiService.Training.getCardSetTrainings(token, slug);
        return trainings;
    },

    async createTraining(slug: string, trainingType: string): Promise<TrainingSummary> {
        const token = localStorage.getItem('accessToken')!;
        return await ApiService.Training.createTraining(token, slug, trainingType);
    }
}