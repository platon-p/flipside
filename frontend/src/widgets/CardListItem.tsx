import { Card } from "@/repository/CardRepository";

interface CardListItemProps {
  card: Card;
}

export function CardListItem({ card }: CardListItemProps) {
  return (
    <div className="flex border-black border-[1px] p-3 gap-4">
      <div className="w-8 flex flex-col justify-center">
        <a>#{card.position}</a>
      </div>
      <div className="flex flex-col gap-4 justify-between w-full">
        <p className="m-0 bg-gray-200">{card.question}</p>
        <p className="m-0 bg-gray-300">{card.answer}</p>
      </div>
    </div>
  );
}
