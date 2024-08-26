import { Block, Input } from "@/shared";

interface CardItemProps {
  position: number;
  question: string;
  answer: string;
  onUpdate: (question: string, answer: string) => void;
}

export function EditableCard({
  position,
  question,
  answer,
  onUpdate,
}: CardItemProps) {
  return (
    <Block className="flex gap-4">
      <div className="flex items-center w-4">#{position}</div>
      <div className="flex flex-col gap-2 justify-between w-full">
        <Input
          placeholder="question"
          value={question}
          onInput={(e) => onUpdate(e.currentTarget.value, answer)}
        />
        <Input
          placeholder="answer"
          value={answer}
          onInput={(e) => onUpdate(question, e.currentTarget.value)}
        />
      </div>
    </Block>
  );
}
