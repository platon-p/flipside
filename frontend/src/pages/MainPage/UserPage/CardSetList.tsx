import { CardSetItem } from "./CardSetItem";
import { CardSet } from "@/repository/CardSetRepository";

interface CardSetListProps {
  cards: CardSet[];
  onClick: (slug: string) => void;
}

export function CardSetList({ cards, onClick }: CardSetListProps) {
  return (
    <div className="flex flex-col gap-4 w-full">
      {cards.map((v, i) => {
        return (
          <CardSetItem
            title={v.title}
            slug={v.slug}
            onClick={() => onClick(v.slug)}
            key={i}
          />
        );
      })}
    </div>
  );
}
