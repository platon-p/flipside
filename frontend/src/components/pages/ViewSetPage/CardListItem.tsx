import { Card } from '@/repository/CardRepository';
import css from './CardListItem.module.css';

export function CardListItem({ card }: { card: Card }) {
    return <div className={css.card}>
        <div className={css.position}>
            <a>#{card.position}</a>
        </div>
        <div className={css.content}>
            <p className={css.question}>{card.question}</p>
            <p className={css.answer}>{card.answer}</p>
        </div>
    </div>

}