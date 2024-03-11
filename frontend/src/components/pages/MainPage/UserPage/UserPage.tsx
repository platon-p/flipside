import { useAuth } from "@/hooks/Auth";
import { CardSet, CardSetRepository } from "@/repository/CardSetRepository";
import { useNavigate } from "react-router-dom";
import { CardSetList } from "./CardSetList";
import { Button } from "@/components/shared/Button";
import "./UserPage.css";
import Navbar from "../../Navbar";
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

    return <div><div className="header">
        <div className="logo"></div>
        <div className="sign-in">
            <a onClick={logout} href="/">выйти</a>
        </div>
    </div>

        <div className="onboarding">
            <h1>Мои наборы</h1>
            {loading && <div>Загрузка...</div>}
            {cardSets && <CardSetList cards={cardSets!} onClick={navigateToCardSet} /> }
            <div className="line"></div>
            <div className="controls">
                <Button className="create-set" onClick={createCardSet}>+ создать новый набор</Button>
            </div>
        </div>

        <Navbar />
    </div>

}