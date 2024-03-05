import { useParams } from "react-router-dom"
import { CardRepository } from "../repository/CardRepository";
import { CardSetRepository } from "../repository/CardSetRepository";

export function CardSet() {
    const { slug } = useParams();
    const cardSet = CardSetRepository.getCardSetBySlug(slug!);
    const cards = CardRepository.getCards(slug!);

    function moveUp(position: number) {
    }

    function moveDown(position: number) {
    }
    if (!cardSet) {
        return <h2>CardSet not found</h2>
    }
    return <div>
        <h2>{cardSet.title}</h2>
        <p>/{cardSet.slug}</p>

        <h4>Cards</h4>
        <div style={{
            display: 'flex',
            flexDirection: 'column',
            gap: '1em'
        }}>
            {cards.map((v, i) => {
                return <div style={{
                    display: 'flex',
                    border: '1px solid black',
                    padding: '0.8em',
                    gap: '1em',
                }} key={i}>
                    <div style={{
                        display: 'flex',
                        flexDirection: 'column',
                    }}>
                        <a onClick={() => moveUp(v.position)}>↑</a>
                        <a>#{v.position}</a>
                        <a onClick={() => moveDown(v.position)}>↓</a>
                    </div>
                    <div style={{
                        display: 'flex',
                        flexDirection: 'column',
                        gap: '1em',
                        width: '100%',
                        justifyContent: 'space-between'
                    }}>
                        <p style={{
                            margin: 0,
                            backgroundColor: '#ddd'
                        }}>{v.question}</p>
                        <p style={{
                            margin: 0,
                            backgroundColor: '#eee'
                        }}>{v.answer}</p>
                    </div>
                </div>
            })}


        </div>
    </div>
}