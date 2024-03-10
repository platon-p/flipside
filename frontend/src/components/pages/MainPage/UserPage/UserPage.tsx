import { useAuth } from "@/hooks/Auth";
import { CardSet, CardSetRepository } from "@/repository/CardSetRepository";
import { useNavigate } from "react-router-dom";
import { CardSetList } from "./CardSetList";
import { Button } from "@/components/shared/Button";
import "./UserPage.css";
import { useEffect, useState } from "react";

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
        <div className="controls">
            <Button onClick={createCardSet}>Create new card set</Button>
            <Button onClick={logout}>Logout</Button>
        </div>
        <h2>Yout sets</h2>
        {loading && <div>Loading...</div>}
        {cardSets && <CardSetList cards={cardSets} onClick={navigateToCardSet} /> }
    </div>
}