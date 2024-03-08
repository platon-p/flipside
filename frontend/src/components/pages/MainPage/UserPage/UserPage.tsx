import { useAuth } from "@/hooks/Auth";
import { CardSetRepository } from "@/repository/CardSetRepository";
import { useNavigate } from "react-router-dom";
import { CardSetList } from "./CardSetList";

export function UserPage() {
    const { userId, logout } = useAuth();
    const navigate = useNavigate();
    const cards = CardSetRepository.getCardSetsByOwner(userId!);

    function navigateToCardSet(slug: string) {
        navigate(`/set/${slug}`)
    }
    
    return <div>
        <p>Hello</p>
        <p onClick={() => { logout() }}>logout</p>
        <CardSetList cards={cards} onClick={navigateToCardSet} />
    </div>
}