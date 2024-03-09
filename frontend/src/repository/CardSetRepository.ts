import { AuthService } from "@/service/AuthService";
import { ApiService } from "../service/ApiService";

export interface CardSet {
    title: string,
    slug: string,
    ownerId: number,
}

export const CardSetRepository = {
    _cardSets: [
        { title: "1. CardSet title", slug: "set1", ownerId: 1 },
        { title: "2. CardSet title", slug: "set2", ownerId: 2 }
    ] as CardSet[],

    async getCardSetBySlug(slug: string): Promise<CardSet> {
        const cardSet = await ApiService.CardSet.getCardSet(slug);
        return {
            title: cardSet.title,
            slug: cardSet.slug,
            ownerId: cardSet.owner_id,
        }
    },

    async createCardSet(title: string, slug: string): Promise<CardSet> {
        const token = AuthService.getToken() ?? '';
        const cardSet = await ApiService.CardSet.createCardSet(token, title, slug);
        return {
            title: cardSet.title,
            slug: cardSet.slug,
            ownerId: cardSet.owner_id,
        }
    },

    async updateCardSet(oldSlug: string, title: string, slug: string): Promise<CardSet> {
        const token = AuthService.getToken() ?? '';
        const cardSet = await ApiService.CardSet.updateCardSet(token, oldSlug, title, slug);
        return {
            title: cardSet.title,
            slug: cardSet.slug,
            ownerId: cardSet.owner_id,
        }
    },

    async deleteCardSet(slug: string): Promise<boolean> {
        const token = AuthService.getToken() ?? '';
        return await ApiService.CardSet.deleteCardSet(token, slug);
    },

    getCardSetsByOwner(ownerId: number): CardSet[] {
        return this._cardSets.filter(it => it.ownerId == ownerId);
    }
}