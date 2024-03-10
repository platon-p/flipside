import { ApiService } from "@/service/ApiService"

export interface Card {
    question: string,
    answer: string,
    position: number,
    owner_id: number,
}

export const CardRepository = {
    async getCards(slug: string): Promise<Card[]> {
        const response = await ApiService.Card.getCards(slug);
        const res = response.map(card => {
            return {
                question: card.question,
                answer: card.answer,
                position: card.position,
                owner_id: card.card_set_id,
            }
        })
        return res;
    }
}