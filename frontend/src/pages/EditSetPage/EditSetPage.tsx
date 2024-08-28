import React, { useEffect } from "react";
import { useParams } from "react-router-dom";
import { Button, Input } from "@/shared";
import { useAuth } from "@/hooks/Auth";
import { EditableCard } from "./EditableCard";
import { useCardSet } from "@/store/cardset";

export default function EditSetPage() {
  const { isAuth } = useAuth();
  const { slug } = useParams();
  const { error, state, addCard } = useCardSet();

  if (!isAuth) {
    return <div>Not authorized</div>;
  }
  if (state === "loading") {
    return <div>Loading...</div>;
  }
  if (state === "error") {
    return <p className="text-red-500">{error}</p>;
  }

  return (
    <div className="max-w-lg mx-auto mt-20">
      <h2 className="text-2xl font-bold">Edit Card Set</h2>
      <MetaDataEditor />
      <CardListEditor />
      <div className="flex gap-2">
        <Button className="w-full" onClick={addCard}>
          Add card
        </Button>
        <Button className="w-full">Submit</Button>
      </div>
    </div>
  );
}

function CardListEditor() {
  const slug = useParams()["slug"]!;
  const { cards, fetchCards } = useCardSet();
  useEffect(() => {
    fetchCards(slug);
  }, [slug]);
  return (
    <>
      <h4 className="text-lg font-medium">Cards</h4>
      <div className="flex flex-col gap-2 mb-2">
        {cards.map((v) => {
          return <EditableCard position={v.position} key={v.position} />;
        })}
      </div>
    </>
  );
}

function MetaDataEditor() {
  const { slug: slugParam } = useParams();
  const { slug, title, fetchSet, setTitle, setSlug } = useCardSet();
  useEffect(() => {
    fetchSet(slugParam!);
  }, [fetchSet, slugParam]);

  const onTitleChange = (title: React.ChangeEvent<HTMLInputElement>) => {
    setTitle(title.currentTarget.value);
  };
  const onSlugChange = (slug: React.ChangeEvent<HTMLInputElement>) => {
    setSlug(slug.currentTarget.value);
  };

  return (
    <div className="w-full flex flex-col gap-2 mt-1">
      <div className="flex items-center gap-2">
        <p className="w-12">Title</p>
        <Input className="w-full" value={title} onChange={onTitleChange} />
      </div>
      <div className="flex items-center gap-2">
        <p className="w-12">Slug</p>
        <Input className="w-full" value={slug} onChange={onSlugChange} />
      </div>
    </div>
  );
}
