import { useAuth } from "@/hooks/Auth";
import { CardSet, CardSetRepository } from "@/repository/CardSetRepository";
import { useNavigate } from "react-router-dom";
import { CardSetList } from "./CardSetList";
import { Button } from "@/components/shared/Button";
import "./UserPage.css";
import { useEffect, useState } from "react";
import Navbar from "../../Navbar";

export function UserPage() {
    const { nickname, logout } = useAuth();
    const navigate = useNavigate();
    const [cardSets, setCardSets] = useState<CardSet[] | undefined>(undefined);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        if (!nickname) return;
        Promise.resolve(CardSetRepository.getCardSetsByOwner(nickname).then(cardSets => {
            setCardSets(cardSets);
            setLoading(false);
        }).catch(e => console.log(e)))
    }, [nickname])
    function navigateToCardSet(slug: string) {
        navigate(`/set/${slug}`)
    }

    function createCardSet() {
        navigate('/create-set');
    }

    return <div>
        <div className="header">
            <div className="logo"></div>
            <div className="sign-in">
                <a onClick={logout} href="/">выйти</a>
            </div>
        </div>

        <div className="onboarding">
            <h1>Мои наборы</h1>
            {loading && <div>Загрузка...</div>}
            {cardSets && <CardSetList cards={cardSets!} onClick={navigateToCardSet} />}
            <div className="line"></div>
            <div className="controls">
                <Button className="create-set" onClick={createCardSet}>+ создать новый набор</Button>
            </div>
        </div>

        <Navbar />
    </div>

}