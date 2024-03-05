import { useParams } from "react-router-dom"

export function CardSet() {
    const { slug } = useParams();
    const cards = [
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
    function moveUp(position: number) {
    }

    function moveDown(position: number) {
    }
    
    return <div>
        <h2>CardTitle</h2>
        <p>/{slug}</p>

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
                        <p style={{ margin: 0 }}>{v.answer}</p>
                    </div>
                </div>
            })}


        </div>
    </div>
}