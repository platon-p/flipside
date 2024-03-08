import { Card } from '@/repository/CardRepository';
import './CardListItem.css';

export function CardListItem({ card }: { card: Card }) {
    return <div className='card' style={{
        display: 'flex',
        border: '1px solid black',
        padding: '0.8em',
        gap: '1em',
    }}>
        <div className='card__position-holder'>
            <a>#{card.position}</a>
        </div>
        <div className='card__content-holder'>
            <p className='card__question'>{card.question}</p>
            <p className='card__answer'>{card.answer}</p>
        </div>
    </div>

}