import './CardSetItem.css';

interface CardSetItemProps {
    title: string
    slug: string
    onClick: (slug: string) => void
}

export function CardSetItem({ title, slug, onClick }: CardSetItemProps) {
    return <div className='card-set-item' onClick={() => onClick(slug)}>
        <p className='card-set-item__title'>{title}</p>
        <p className='card-set-item__slug'>/{slug}</p>
    </div>
}