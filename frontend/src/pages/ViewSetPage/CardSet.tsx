import { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { CardSet, CardSetRepository } from "@/repository/CardSetRepository";
import { Card, CardRepository } from "@/repository/CardRepository";
import { useAuth } from "@/hooks/Auth";
import {
  TrainingRepository,
  TrainingSummary,
} from "@/repository/TrainingRepository";
import { CardSetControls, TrainingsWidget, CardListWidget } from "@/widgets";
import { Block } from "@/shared";

export function ViewSetPage() {
  const { userId } = useAuth();
  const { slug } = useParams();
  const navigate = useNavigate();

  const [cardSet, setCardSet] = useState<CardSet | undefined>();
  const [cards, setCards] = useState<Card[] | undefined>();
  const [trainings, setTrainings] = useState<TrainingSummary[] | undefined>();

  const [loading, setLoading] = useState(true);
  const [errorMessage, setErrorMessage] = useState<string | null>(null);

  useEffect(() => {
    async function loadCardSet() {
      try {
        const cardSet = await CardSetRepository.getCardSetBySlug(slug!);
        setCardSet(cardSet);
      } catch (e) {
        if (e === "Card Set not found") {
          console.log("Card Set not found");
        } else if (typeof e === "string") {
          setErrorMessage(e);
        } else {
          console.error(e);
        }
        setLoading(false);
      }
    }

    async function loadCards() {
      try {
        const cards = await CardRepository.getCards(slug!);
        setCards(cards);
      } catch (e) {
        console.error(e);
      }
    }

    async function loadTrainings() {
      try {
        const trainings = await TrainingRepository.getCardSetTrainings(slug!);
        setTrainings(trainings);
      } catch (e) {
        console.error(e);
      }
    }
    Promise.all([loadCardSet(), loadCards(), loadTrainings()]).then(() => {
      setLoading(false);
    });
  }, [slug]);

  const goHome = () => navigate("/");
  const edit = () => navigate("edit");

  function remove() {
    CardSetRepository.deleteCardSet(slug!)
      .then((res) => {
        console.log("CardSet deleted", res);
        navigate("/");
      })
      .catch((e) => {
        console.log(e);
      });
  }

  function createTraining(trainingType: string) {
    TrainingRepository.createTraining(slug!, trainingType)
      .then((training) => {
        console.log("Training created", training);
        window.location.reload();
      })
      .catch((e) => {
        console.log(e);
      });
  }

  if (errorMessage) {
    return <p className="text-red-500">{errorMessage}</p>;
  }
  if (loading) {
    return <h2>Loading...</h2>;
  }
  if (!cardSet) {
    return <h2>CardSet not found</h2>;
  }
  return (
    <div className="flex flex-col gap-2 p-2 max-w-md m-auto">
      <p onClick={goHome}>{"<"} home</p>
      <Block className="flex flex-col gap-2">
        <div className="flex items-center justify-between">
          <h2 className="text-2xl font-semibold">{cardSet.title}</h2>
          <p className="text-gray-600">/{cardSet.slug}</p>
        </div>
        {cardSet.ownerId === userId && (
          <CardSetControls edit={edit} remove={remove} />
        )}
      </Block>
      <TrainingsWidget trainings={trainings} />
      <CardListWidget cards={cards} />
    </div>
  );
}
