import { Block, Input } from "@/shared";
import { useCardSet } from "@/store/cardset";

interface EditableCardProps {
  position: number;
}

export function EditableCard({ position }: EditableCardProps) {
  const { selectCard, editCard } = useCardSet();
  const card = selectCard(position);
  const editQuestion = (e: React.FormEvent<HTMLInputElement>) => {
    editCard({
      position,
      question: e.currentTarget.value,
      answer: card.answer,
    });
  };
  const editAnswer = (e: React.FormEvent<HTMLInputElement>) => {
    editCard({
      position,
      question: card.question,
      answer: e.currentTarget.value,
    });
  };
  return (
    <Block className="relative flex gap-4 hover:[&>.zov]:bg-red-400">
      <div
        className={`zov -right-3 -top-3 absolute
        w-8 h-8 rounded-full bg-orange-500
        cursor-pointer text-white
        flex items-center justify-center
        text-xl font-medium
        `}
      >
        {"-"}
      </div>
      <div className="flex items-center w-4">#{position}</div>
      <div className="flex flex-col gap-2 justify-between w-full">
        <Input
          placeholder="question"
          value={card.question}
          onInput={editQuestion}
        />
        <Input placeholder="answer" value={card.answer} onInput={editAnswer} />
      </div>
    </Block>
  );
}
