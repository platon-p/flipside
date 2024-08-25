import { Button } from "@/shared";

interface CardSetControlProps {
  edit: () => void;
  remove: () => void;
}

export function CardSetControls({ edit, remove }: CardSetControlProps) {
  return (
    <div className="flex gap-2">
      <Button className="w-full" onClick={edit}>
        Edit
      </Button>
      <Button className="w-full" onClick={remove}>
        Delete
      </Button>
    </div>
  );
}
