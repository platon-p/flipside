import { client } from "./client";
import { CardSet } from "@/repository/CardSetRepository";

export const ProfileApi = {
  async getUserCardSets(nickname: string): Promise<CardSet[]> {
    const response = await client.get(`api/users/${nickname}/sets`); // TODO: url
    return await response.json();
  },

  async getUserProfile(nickname: string): Promise<any> {
    const resp = await client.get(`api/users/${nickname}/profile`);
    return await resp.json();
  }
};
