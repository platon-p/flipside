import { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { Card, CardRepository } from "@/repository/CardRepository";
import { CardSet, CardSetRepository } from "@/repository/CardSetRepository";
import { Button, Input } from "@/shared";
import { useAuth } from "@/hooks/Auth";
import { EditableCard } from "./EditableCard";
// import { RootState } from "@/reducers/store";

// const { editCurrent } = cardSetSlice.actions;

export default function EditSetPage() {
  const { isAuth, userId } = useAuth();
  const { slug } = useParams();
  const navigate = useNavigate();
  // const dispatch = useDispatch();

  // const state = useSelector((state: RootState) => state.cardSet);

  const [cardSet, setCardSet] = useState<CardSet | undefined>();
  const [displayCards, setDisplayCards] = useState<Card[] | undefined>();
  const [loading, setLoading] = useState(true);

  const [errorMessage, setErrorMessage] = useState<string | undefined>();
  const [createdCards, setCreatedCards] = useState(new Array<Card>());

  useEffect(() => {
    async function loadCardSet(): Promise<void> {
      try {
        const i = await CardSetRepository.getCardSetBySlug(slug!);
        setCardSet(i);
        // dispatch(editCurrent({ title: i.title, slug: i.slug }));
      } catch (e) {
        setErrorMessage(e?.toString());
      }
    }
    async function loadCards(): Promise<void> {
      const i = await CardRepository.getCards(slug!);
      setDisplayCards(i);
    }
    Promise.all([loadCardSet(), loadCards()]).then(() => setLoading(false));
  }, [slug]);

  const handleUpdate = (position: number, question: string, answer: string) => {
    if (!displayCards) return;
    displayCards[position].question = question;
    displayCards[position].answer = answer;
    setDisplayCards([...displayCards]);
  };

  async function submit() {
    try {
      await CardRepository.createCards(slug!, createdCards);
    } catch (e) {
      console.log(e);
      return;
    }

    CardSetRepository.updateCardSet(
      slug!,
      'state.current.title',
      'state.current.slug'
    )
      .then((res) => {
        navigate(`/set/${res.slug}`);
        console.log("CardSet updated", res);
      })
      .catch((e) => {
        console.log(e);
      });
  }

  function addCard() {
    if (!displayCards) return;
    displayCards.push({
      question: "question",
      answer: "answer",
      position: (displayCards.at(-1)?.position ?? 0) + 1,
    });
    setDisplayCards([...displayCards]);
    setCreatedCards([...createdCards, displayCards.at(-1)!]);
  }

  if (!isAuth) {
    return <div>Not authorized</div>;
  }
  if (loading) {
    return <div>Loading...</div>;
  }
  if (errorMessage) {
    return <p style={{ color: "red" }}>{errorMessage}</p>;
  }
  if (cardSet === undefined) {
    return <div>Card Set not found</div>;
  }

  if (cardSet.ownerId !== userId) {
    return <div>Not authorized</div>;
  }

  return (
    <div className="max-w-lg mx-auto mt-20">
      <h2 className="text-2xl font-bold">Edit Card Set</h2>
      <MetaDataEditor />
      <Button onClick={submit}>Submit</Button>
      <h4 className="text-lg font-medium">Cards</h4>
      <div className="flex flex-col gap-2 mb-2">
        {displayCards?.map((v, i) => {
          return (
            <EditableCard
              position={v.position}
              question={v.question}
              answer={v.answer}
              onUpdate={(q, a) => handleUpdate(i, q, a)}
              key={v.position}
            />
          );
        })}
      </div>
      <Button onClick={addCard}>Add card</Button>
    </div>
  );
}

function MetaDataEditor() {
  // const state = useSelector((state: RootState) => state.cardSet);

  return (
    <div className="w-full flex flex-col gap-2 mt-1">
      <div className="flex items-center gap-2">
        <p className="w-12">Title</p>
        <Input
          className="w-full"
          // value={state.current.title}
          // onChange={(e) => dispatch(setTitle(e.currentTarget.value))}
        />
      </div>
      <div className="flex items-center gap-2">
        <p className="w-12">Slug</p>
        <Input
          className="w-full"
          // value={state.current.slug}
          // onChange={(e) => dispatch(setSlug(e.currentTarget.value))}
        />
      </div>
    </div>
  );
}
