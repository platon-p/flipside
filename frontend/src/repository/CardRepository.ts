import { ApiService } from "@/service/ApiService/ApiService"

export interface Card {
    question: string,
    answer: string,
    position: number,
}

export const CardRepository = {
    async getCards(slug: string): Promise<Card[]> {
        const response = await ApiService.Card.getCards(slug);
        const res = response.map(card => {
            return {
                question: card.question,
                answer: card.answer,
                position: card.position,
            }
        })
        return res;
    },

    async createCards(slug: string, cards: Card[]): Promise<Card[]> {
        const token = localStorage.getItem('accessToken')!;
        const response = await ApiService.Card.createCards(token, slug, cards.map(card => {
            return {
                question: card.question,
                answer: card.answer,
                position: card.position,
            }
        }));
        return response.map(card => {
            return {
                question: card.question,
                answer: card.answer,
                position: card.position,
            }
        });
    },
}