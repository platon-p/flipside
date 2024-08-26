import { client, config } from "./client";

export interface CardRequest {
  question: string;
  answer: string;
  position: number;
}

export interface CardResponse {
  question: string;
  answer: string;
  position: number;
  card_set_id: number;
}

export const CardApi = {
  async createCards(
    token: string,
    slug: string,
    request: CardRequest[],
  ): Promise<CardResponse[]> {
    const response = await client.post(`${config.cards}/${slug}`, {
      headers: { Authorization: `Bearer ${token}` },
      json: request,
    });
    return await response.json();
  },

  async getCards(slug: string): Promise<CardResponse[]> {
    const response = await client.get(`${config.cards}/${slug}`);
    return await response.json();
  },

  async updateCards(
    token: string,
    slug: string,
    request: CardRequest[],
  ): Promise<CardResponse[]> {
    const response = await client.put(`${config.cards}/${slug}`, {
      headers: { Authorization: `Bearer ${token}` },
      json: request,
    });
    return await response.json();
  },

  async deleteCards(
    token: string,
    slug: string,
    positions: number[],
  ): Promise<void> {
    await client.delete(`${config.cards}/${slug}`, {
      headers: { Authorization: `Bearer ${token}` },
      json: positions,
    });
  },
};
