import { config, client } from "./client";

export interface CardSetResponse {
  title: string;
  slug: string;
  owner_id: number;
}

export const CardSetApi = {
  async createCardSet(
    token: string,
    title: string,
    slug: string,
  ): Promise<CardSetResponse> {
    const response = await client.post(`${config.cardSet}/`, {
      headers: { Authorization: `Bearer ${token}` },
      json: { title: title, slug: slug },
    });
    return await response.json();
  },

  async getCardSet(slug: string): Promise<CardSetResponse> {
    const response = await client.get(`${config.cardSet}/${slug}`);
    return await response.json();
  },

  async updateCardSet(
    token: string,
    slug: string,
    newTitle: string,
    newSlug: string,
  ): Promise<CardSetResponse> {
    const response = await client.put(`${config.cardSet}/${slug}`, {
      headers: { Authorization: `Bearer ${token}` },
      json: { title: newTitle, slug: newSlug },
    });
    return await response.json();
  },

  async deleteCardSet(token: string, slug: string): Promise<void> {
    await client.delete(`${config.cardSet}/${slug}`, {
      headers: { Authorization: `Bearer ${token}` },
    });
  },
};
