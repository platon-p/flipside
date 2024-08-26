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
  if (loading) {
    return <p>Загрузка...</p>;
  }
  return (
    <div className="flex flex-col mx-auto max-w-xs px-2 gap-2">
      <h1 className="text-2xl font-bold">Мои наборы</h1>
      <a className="text-orange-500" onClick={logout} href="/">
        выйти
      </a>
      {cardSets && (
        <CardSetList cards={cardSets!} onClick={navigateToCardSet} />
      )}
      <Button className="w-full" onClick={createCardSet}>
        + создать новый набор
      </Button>
    </div>
  );
}
