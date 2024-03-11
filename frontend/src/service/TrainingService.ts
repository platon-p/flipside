import { ApiService, TaskResult, TrainingTask } from "./ApiService";

export const TrainingService = {
    getNextTask(trainingId: number): Promise<TrainingTask> {
        const token = localStorage.getItem('accessToken')!;
        return ApiService.Training.getNextTask(token, trainingId);
    },

    submitAnswer(trainingId: number, answer: string): Promise<TaskResult> {
        const token = localStorage.getItem('accessToken')!;
        return ApiService.Training.submitAnswer(token, trainingId, answer);
    }
}