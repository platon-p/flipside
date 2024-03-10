import { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom"
import { CardSet as CardSetModel, CardSetRepository } from "@/repository/CardSetRepository";
import { Card, CardRepository } from "@/repository/CardRepository";
import { useAuth } from "@/hooks/Auth";
import { Button } from "@/components/shared/Button";
import { CardListItem } from "./CardListItem";
import { TrainingRepository, TrainingSummary } from "@/repository/TrainingRepository";

export function ViewSetPage() {
    const { userId } = useAuth();
    const { slug } = useParams();
    const navigate = useNavigate();

    const [cardSet, setCardSet] = useState<CardSetModel | undefined>();
    const [cards, setCards] = useState<Card[] | undefined>();
    const [trainings, setTrainings] = useState<TrainingSummary[] | undefined>();

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

        async function loadTrainings() {
            try {
                const response = await TrainingRepository.getCardSetTrainings(slug!);
                setTrainings(response);
            } catch (e) {
                setErrorMessage('Failed to load trainings');
            }
        }

        Promise.all([loadCardSet(), loadCards(), loadTrainings()]);
    }, [slug])

    function goHome() {
        navigate('/');
    }

    function edit() {
        navigate('edit')
    }

    function remove() {
        CardSetRepository.deleteCardSet(slug!)
            .then(res => {
                console.log('CardSet deleted', res)
                navigate('/')
            })
            .catch(e => {
                console.log(e)
            })
    }

    async function createBasicTraining() {
        try {
            const res = await TrainingRepository.createTraining(slug!, 'basic')
            console.log('ok', res);
        } catch (e) {
            console.error(e);
        }
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
            <Button onClick={edit}>Edit</Button>
            <Button onClick={remove}>Delete</Button>
        </div>}

        <h4>Trainings</h4>
        <div>
            <Button onClick={createBasicTraining}>Create basic training</Button>
        </div>
        {trainings?.length === 0 ? <p>Empty list</p> : <div style={{
            display: 'flex',
            flexDirection: 'column',
            gap: '1em'
        }}>
            {trainings?.map((v, i) => <div key={i} style={{
                display: 'flex',
                justifyContent: 'space-around',
                backgroundColor: 'lightgray',
                alignItems: 'center'
            }}>
                <p>{v.training_type}</p>
                <p style={{
                    backgroundColor: '#0b0',
                    color: '#141',
                    padding: '0.3em'
                }}>{v.status}</p>
                <p style={{ color: "green" }}>+{v.count_right}</p>
                <p style={{ color: "red" }}>-{v.count_wrong}</p>
                <Button>Train</Button>
            </div>)}
        </div>
        }


        <h4>Cards</h4>
        {cards?.length === 0 ? <p>Empty list</p> : <div style={{
            display: 'flex',
            flexDirection: 'column',
            gap: '1em'
        }}>
            {cards?.map((v, i) => <CardListItem card={v} key={i} />)}
        </div>
        }
    </div>
}