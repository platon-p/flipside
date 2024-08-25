import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { useAuth } from "@/hooks/Auth";
import { CardSet, CardSetRepository } from "@/repository/CardSetRepository";
import { Button } from "@/shared";
import { CardSetList } from "./CardSetList";

export function UserPage() {
  const { logout, nickname } = useAuth();
  const navigate = useNavigate();
  const [cardSets, setCardSets] = useState<CardSet[] | undefined>();
  const [loading, setLoading] = useState(true);

  function navigateToCardSet(slug: string) {
    navigate(`/set/${slug}`);
  }

  function createCardSet() {
    navigate("/create-set");
  }

  useEffect(() => {
    CardSetRepository.getCardSetsByOwner(nickname!).then((cardSets) => {
      setCardSets(cardSets);
      setLoading(false);
    });
  }, [nickname]);

  return (
    <div>
      <div className="header">
        <h2>Мои наборы</h2>
        <div className="logo"></div>
        <div className="sign-in">
          <a className="text-orange-500" onClick={logout} href="/">
            выйти
          </a>
        </div>
      </div>
      {loading && <div>Загрузка...</div>}
      <div
        className="flex flex-col items-center mx-auto"
        style={{ width: "80%" }}
      >
        {cardSets && (
          <CardSetList cards={cardSets!} onClick={navigateToCardSet} />
        )}
        <div className="controls w-full">
          <Button onClick={createCardSet}>+ создать новый набор</Button>
        </div>
      </div>
    </div>
  );
}
