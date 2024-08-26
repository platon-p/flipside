import { Card } from "@/repository/CardRepository";
import { CardListItem } from "./CardListItem";

export interface CardListWidgetProps {
  cards: Card[] | undefined;
}

export function CardListWidget({ cards }: CardListWidgetProps) {
  return (
    <div className="flex flex-col">
      <h4 className="text-lg font-semibold">Cards</h4>
      {!cards || cards.length === 0 ? (
        <p>Empty list</p>
      ) : (
        <div className="flex flex-col gap-4">
          {cards?.map((v, i) => (
            <CardListItem card={v} key={i} />
          ))}
        </div>
      )}
    </div>
  );
}
