import { useNavigate, useParams } from "react-router-dom"
import { CardSet as CardSetModel, CardSetRepository } from "../repository/CardSetRepository";
import { useEffect, useState } from "react";
import { Card, CardRepository } from "@/repository/CardRepository";
import { useAuth } from "@/hooks/Auth";
import { Button } from "@/components/Button";

export function CardSet() {
    const { userId } = useAuth();
    const { slug } = useParams();
    const navigate = useNavigate();

    const [cardSet, setCardSet] = useState<CardSetModel | undefined>();
    const [cards, setCards] = useState<Card[] | undefined>();

    const [loading, setLoading] = useState(true);
    const [errorMessage, setErrorMessage] = useState<string | null>(null);

    useEffect(() => {
        async function loadCardSet() {
            try {
                const cardSet = await CardSetRepository.getCardSetBySlug(slug!);
                setCardSet(cardSet);
                setLoading(false);
            } catch (e) {
                if (e === 'Card Set not found') {
                    console.log('Card Set not found');
                } else if (typeof e === 'string') {
                    setErrorMessage(e);
                } else {
                    console.error(e);
                }
                setLoading(false);
            }
        }

        async function loadCards() {
            try {
                const cards = await CardRepository.getCards(slug!);
                setCards(cards);
            } catch (e) {
                console.error(e);
            }
        }

        Promise.all([loadCardSet(), loadCards()]);
    }, [slug])

    function goHome() {
        navigate('/');
    }

    if (errorMessage) {
        return <p style={{
            color: 'red'
        }}>{errorMessage}</p>
    }
    if (loading) {
        return <h2>Loading...</h2>
    }

    if (!cardSet) {
        return <h2>CardSet not found</h2>
    }

    return <div>
        <p onClick={goHome}>home</p>
        <h2>{cardSet.title}</h2>
        <p>/{cardSet.slug}</p>

        {cardSet.ownerId === userId && <div>
            <Button>Edit</Button>
            <Button>Delete</Button>
        </div>}

        <h4>Cards</h4>
        <div style={{
            display: 'flex',
            flexDirection: 'column',
            gap: '1em'
        }}>
            {cards && cards.map((v, i) => {
                return <div style={{
                    display: 'flex',
                    border: '1px solid black',
                    padding: '0.8em',
                    gap: '1em',
                }} key={i}>
                    <div style={{
                        display: 'flex',
                        flexDirection: 'column',
                        justifyContent: 'center',
                    }}>
                        <a>#{v.position}</a>
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