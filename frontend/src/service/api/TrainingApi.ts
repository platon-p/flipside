import { client, config } from "./client";

export interface TrainingSummaryResponse {
  id: number;
  card_set_id: number;
  training_type: string;
  status: string;
  count_right: number;
  count_wrong: number;
}

export interface TrainingTask {
  question: string;
  question_type: string;
  answers: string[];
}

export interface TaskResult {
  answer: string;
  is_correct: boolean;
}

export const TrainingApi = {
  async getCardSetTrainings(
    token: string,
    slug: string
  ): Promise<TrainingSummaryResponse[]> {
    const response = await client.get(`${config.cardSet}/${slug}/trainings`, {
      headers: { Authorization: `Bearer ${token}` },
    });
    return await response.json();
  },

  async createTraining(
    token: string,
    slug: string,
    trainingType: string
  ): Promise<TrainingSummaryResponse> {
    const response = await client.post(`${config.cardSet}/${slug}/trainings`, {
      searchParams: { type: trainingType },
      headers: { Authorization: `Bearer ${token}` },
    });
    return await response.json();
  },

  async getNextTask(token: string, trainingId: number): Promise<TrainingTask> {
    const response = await client.get(`${config.training}/${trainingId}/next`, {
      headers: { Authorization: `Bearer ${token}` },
    });
    return await response.json();
  },

  async submitAnswer(
    token: string,
    trainingId: number,
    answer: string
  ): Promise<TaskResult> {
    const response = await client.post(
      `${config.training}${trainingId}/submit`,
      {
        searchParams: { answer: answer },
        headers: { Authorization: `Bearer ${token}` },
      }
    );
    return await response.json();
  },
};
