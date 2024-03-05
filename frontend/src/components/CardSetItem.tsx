interface CardSetItemProps {
    title: string
    slug: string
}

export function CardSetItem({title, slug}: CardSetItemProps) {
    return <div style={{
        border: 'solid black',
        display: 'flex',
        flexDirection: 'column',
        padding: '1em',
        gap: '1em',
        backgroundColor: '#eee'
    }}>
        <p style={{
            display: 'inline-block',
            margin: 0,
            fontSize: '1.3em',
            fontWeight: 'bold',
        }}>{title}</p>
        <p style={{
            display: 'inline-block',
            margin: 0
        }}>/{slug}</p>
    </div>
}