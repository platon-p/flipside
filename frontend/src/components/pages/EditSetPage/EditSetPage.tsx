import { Button } from "@/components/shared/Button";
import { Input } from "@/components/shared/Input";
import { useAuth } from "@/hooks/Auth";
import { Card, CardRepository } from "@/repository/CardRepository";
import { CardSet, CardSetRepository } from "@/repository/CardSetRepository";
import { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom"
import { CardItem } from "./CardItem";

export default function EditSetPage() {
    const { isAuth, userId } = useAuth();
    const { slug: slugParam } = useParams();
    const navigate = useNavigate();

    const [cardSet, setCardSet] = useState<CardSet | undefined>();
    const [initCards, setInitCards] = useState<Card[] | undefined>();
    const [displayCards, setDisplayCards] = useState<Card[] | undefined>();
    const [loading, setLoading] = useState(true);

    const [title, setTitle] = useState<string>('');
    const [slug, setSlug] = useState<string>('');
    const [errorMessage, setErrorMessage] = useState<string | undefined>();

    const [createdCards, setCreatedCards] = useState(new Array<Card>());

    useEffect(() => {
        function loadCardSet() {
            CardSetRepository.getCardSetBySlug(slugParam!)
                .then(i => {
                    setCardSet(i)
                    setTitle(i.title)
                    setSlug(i.slug)
                })
                .catch(e => {
                    setErrorMessage(e?.toString())
                })
        }
        function loadCards() {
            CardRepository.getCards(slugParam!)
                .then(i => {
                    setInitCards(i)
                    setDisplayCards(i)
                })
        }
        Promise.all([loadCardSet(), loadCards()]).then(() => setLoading(false))
    }, [slugParam])

    const handleUpdate = (position: number, question: string, answer: string) => {
        if (!displayCards) return;
        displayCards[position].question = question
        displayCards[position].answer = answer
        setDisplayCards([...displayCards])
    }

    async function submit() {
        try {
            await CardRepository.createCards(slugParam!, createdCards)
        } catch (e) {
            console.log(e)
            return
        }

        CardSetRepository.updateCardSet(slugParam!, title, slug)
            .then(res => {
                navigate(`/set/${res.slug}`)
                console.log('CardSet updated', res)
            })
            .catch(e => {
                console.log(e)
            })
    }

    function addCard() {
        if (!displayCards) return;
        displayCards.push({
            question: 'question',
            answer: 'answer',
            position: (displayCards.at(-1)?.position ?? 0) + 1,
        })
        setDisplayCards([...displayCards])
        setCreatedCards([...createdCards, displayCards.at(-1)!])
    }

    if (!isAuth) {
        return <div>Not authorized</div>
    }
    if (loading) {
        return <div>Loading...</div>
    }
    if (errorMessage) {
        return <p style={{
            color: 'red',
        }}>{errorMessage}</p>
    }
    if (cardSet === undefined) {
        return <div>Card Set not found</div>
    }

    if (cardSet.ownerId !== userId) {
        return <div>Not authorized</div>
    }

    return <div>
        <h2>Edit Card Set</h2>
        <div className="cardset-data">
            <div style={{
                display: 'flex',
                alignItems: 'center',
                gap: '10px'
            }}>
                <p style={{ margin: 0 }}>Title</p>
                <Input value={title} onChange={e => setTitle(e.currentTarget.value)} />
            </div>
            <div style={{
                display: 'flex',
                alignItems: 'center',
                gap: '10px'
            }}>
                <p>Slug</p>
                <Input value={slug} onChange={e => setSlug(e.currentTarget.value)} />
            </div>
            <Button onClick={submit}>Submit</Button>
        </div>

        <h4>Cards</h4>
        <div style={{
            display: 'flex',
            flexDirection: 'column',
            gap: '1em',
            marginBottom: '1em'
        }}>
            {displayCards?.map((v, i) => {
                return <CardItem
                    position={v.position}
                    question={v.question}
                    answer={v.answer}
                    onUpdate={(q, a) => handleUpdate(i, q, a)}
                    key={v.position}
                />
            })}
        </div>
        <Button onClick={addCard}>Add card</Button>
    </div>
}