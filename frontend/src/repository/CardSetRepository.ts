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
    getCardSetBySlug(slug: string): CardSet | undefined {
        return this._cardSets.find(it => it.slug == slug);
    },
    getCardSetsByOwner(ownerId: number): CardSet[] {
        return this._cardSets.filter(it => it.ownerId == ownerId);
    }
}