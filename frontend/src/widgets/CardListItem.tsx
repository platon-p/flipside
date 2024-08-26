import { Card } from "@/repository/CardRepository";
import { Block } from "@/shared";

interface CardListItemProps {
  card: Card;
}

export function CardListItem({ card }: CardListItemProps) {
  return (
    <Block className="flex p-3 gap-2">
      <div className="w-8 flex flex-col justify-center">
        <p>#{card.position}</p>
      </div>
      <div className="flex flex-col gap-3 justify-between w-full">
        <p className="">{card.question}</p>
        <p className="">{card.answer}</p>
      </div>
    </Block>
  );
}
