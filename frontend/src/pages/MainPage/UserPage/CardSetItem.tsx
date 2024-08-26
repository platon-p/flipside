import { Block } from "@/shared";

interface CardSetItemProps {
  title: string;
  slug: string;
  onClick: () => void;
}

export function CardSetItem({ title, slug, onClick }: CardSetItemProps) {
  return (
    <Block className="flex flex-col cursor-pointer" onClick={onClick}>
      <p className="font-semibold text-lg">{title}</p>
      <p>/{slug}</p>
    </Block>
  );
}
