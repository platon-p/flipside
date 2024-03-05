export interface Card {
    question: string,
    answer: string,
    position: number,
}

export const CardRepository = {
    getCards(slug: string): Card[] {
        return [
            {
                question: '1. Card question',
                answer: '1. Card answer',
                position: 1
            },
            {
                question: '2. Card question',
                answer: '2. Card answer',
                position: 2
            }
        ]
    }
}