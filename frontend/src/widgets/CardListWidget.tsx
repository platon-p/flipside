import { Card } from "@/repository/CardRepository";
import { Block } from "@/shared";
import { CardListItem } from "./CardListItem";

export interface CardListWidgetProps {
  cards: Card[] | undefined;
}

export function CardListWidget({ cards }: CardListWidgetProps) {
  return (
    <Block>
      <h4 className="text-xl font-semibold">Cards</h4>
      {cards?.length === 0 ? (
        <p>Empty list</p>
      ) : (
        <div className="flex flex-col gap-4">
          {cards?.map((v, i) => <CardListItem card={v} key={i} />)}
        </div>
      )}
    </Block>
  );
}
