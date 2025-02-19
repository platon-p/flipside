import { AuthService } from "@/service/AuthService";
import { ApiService } from "@/service/ApiService";

export interface CardSet {
  title: string;
  slug: string;
  ownerId: number;
}

export const CardSetRepository = {
  async getCardSetBySlug(slug: string): Promise<CardSet> {
    const cardSet = await ApiService.CardSet.getCardSet(slug);
    return {
      title: cardSet.title,
      slug: cardSet.slug,
      ownerId: cardSet.owner_id,
    };
  },

  async createCardSet(title: string, slug: string): Promise<CardSet> {
    const token = AuthService.getToken() ?? "";
    const cardSet = await ApiService.CardSet.createCardSet(token, title, slug);
    return {
      title: cardSet.title,
      slug: cardSet.slug,
      ownerId: cardSet.owner_id,
    };
  },

  async updateCardSet(
    oldSlug: string,
    title: string,
    slug: string,
  ): Promise<CardSet> {
    const token = AuthService.getToken() ?? "";
    const cardSet = await ApiService.CardSet.updateCardSet(
      token,
      oldSlug,
      title,
      slug,
    );
    return {
      title: cardSet.title,
      slug: cardSet.slug,
      ownerId: cardSet.owner_id,
    };
  },

  async deleteCardSet(slug: string): Promise<void> {
    const token = AuthService.getToken() ?? "";
    await ApiService.CardSet.deleteCardSet(token, slug);
  },

  async getCardSetsByOwner(nickname: string): Promise<CardSet[]> {
    return (await ApiService.Profile.getUserCardSets(nickname)).map(
      (cardSet) => {
        return {
          title: cardSet.title,
          slug: cardSet.slug,
          ownerId: cardSet.ownerId,
        };
      },
    );
  },
};
