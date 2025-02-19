import { TrainingSummary } from "@/repository/TrainingRepository";
import { Block, Button } from "@/shared";

interface TrainingItemProps {
  training: TrainingSummary;
  onClick: () => void;
}

export function TrainingItem({ training, onClick }: TrainingItemProps) {
  return (
    <Block className="flex justify-around w-full bg-gray-300 items-center">
      <p>{training.id}</p>
      <p>{training.status}</p>
      <p>{training.training_type}</p>
      <p className="text-green-600">+{training.count_right}</p>
      <p className="text-red-600">-{training.count_wrong}</p>
      <Button onClick={onClick}>Start</Button>
    </Block>
  );
}
