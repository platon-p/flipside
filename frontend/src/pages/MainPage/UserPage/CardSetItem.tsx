interface CardSetItemProps {
  title: string;
  slug: string;
  onClick: () => void;
}

export function CardSetItem({ title, slug, onClick }: CardSetItemProps) {
  return (
    <div
      className="border border-black flex flex-col px-3 py-2 gap-2 bg-gray-200"
      onClick={onClick}
    >
      <p className="inline-block font-bold text-lg">{title}</p>
      <p className="inline-block">/{slug}</p>
    </div>
  );
}
