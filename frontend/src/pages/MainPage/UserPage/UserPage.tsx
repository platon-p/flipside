import { useAuth } from "@/hooks/Auth";
import { CardSet, CardSetRepository } from "@/repository/CardSetRepository";
import { useNavigate } from "react-router-dom";
import { CardSetList } from "./CardSetList";
import { Button } from "@/shared/Button";
import "./UserPage.css";
import { useEffect, useState } from "react";

export function UserPage() {
    const { logout, nickname } = useAuth();
    const navigate = useNavigate();
    const [cardSets, setCardSets] = useState<CardSet[] | undefined>();
    const [loading, setLoading] = useState(true);


    function navigateToCardSet(slug: string) {
        navigate(`/set/${slug}`)
    }

    function createCardSet() {
        navigate('/create-set');
    }

    useEffect(() => {
        CardSetRepository.getCardSetsByOwner(nickname!)
            .then(cardSets => {
                setCardSets(cardSets);
                setLoading(false);
            })
    }, [nickname])

    return <div>
        <div className="header">
            <h2>Мои наборы</h2>
            <div className="logo"></div>
            <div className="sign-in">
                <a style={{ color: '#F1694F' }} onClick={logout} href="/">выйти</a>
            </div>
        </div>
        {loading && <div>Загрузка...</div>}
        <div style={{
            width: '80%',
            display: 'flex',
            flexDirection: 'column',
            alignItems: 'center',
            margin: '0 auto',
        }}>

            {cardSets && <CardSetList cards={cardSets!} onClick={navigateToCardSet} />}
            <div style={{width: '100%'}} className="controls">
                <Button className="create-set" onClick={createCardSet}>+ создать новый набор</Button>
            </div>
        </div>
    </div>

}