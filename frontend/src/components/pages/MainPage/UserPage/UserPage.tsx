import { useAuth } from "@/hooks/Auth";
import { CardSetRepository } from "@/repository/CardSetRepository";
import { useNavigate } from "react-router-dom";
import { CardSetList } from "./CardSetList";
import { Button } from "@/components/shared/Button";
import "./UserPage.css";

export function UserPage() {
    const { userId, logout } = useAuth();
    const navigate = useNavigate();
    const cards = CardSetRepository.getCardSetsByOwner(userId!);

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
        <CardSetList cards={cards} onClick={navigateToCardSet} />
    </div>
}