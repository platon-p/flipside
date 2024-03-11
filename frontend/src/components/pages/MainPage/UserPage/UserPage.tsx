import { useAuth } from "@/hooks/Auth";
import { CardSetRepository } from "@/repository/CardSetRepository";
import { useNavigate } from "react-router-dom";
import { CardSetList } from "./CardSetList";
import { Button } from "@/components/shared/Button";
import "./UserPage.css";
import Navbar from "../../Navbar";

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

    return <div><div className="header">
    <div className="logo"></div>
    <div className="sign-in">
        <a onClick={logout} href="/">выйти</a>
    </div>
</div>

<div className="onboarding">
    <h1>Мои наборы</h1>
    <CardSetList cards={cards} onClick={navigateToCardSet} />
    <div className="line"></div>
    <div className="controls">
            <Button className="create-set" onClick={createCardSet}>+ создать новый набор</Button>
        </div>
    </div>
        
        <Navbar/>
    </div>
    
}