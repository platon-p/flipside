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
    getCardSetsByOwner(ownerId: number): CardSet[] {
        return this._cardSets.filter(it => it.ownerId == ownerId);
    }
}