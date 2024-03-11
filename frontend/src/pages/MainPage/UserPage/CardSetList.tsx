import { CardSetItem } from './CardSetItem';
import { CardSet } from "@/repository/CardSetRepository";
import './CardSetList.css';

export function CardSetList({ cards, onClick }: {
    cards: CardSet[],
    onClick: (slug: string) => void,
}) {
    return <div className="card-set-list" style={{
        width: '100%',
    }}>
        {cards.map((v, i) => {
            return <CardSetItem title={v.title} slug={v.slug} onClick={onClick} key={i} />
        })}
    </div >
}