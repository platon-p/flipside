import { useNavigate } from "react-router-dom";
import { TrainingSummary } from "@/repository/TrainingRepository";
import { Block, Button } from "@/shared";
import { TrainingItem } from "./TrainingItem";

interface TrainingsWidgetProps {
  trainings: TrainingSummary[] | undefined;
}

export function TrainingsWidget({ trainings }: TrainingsWidgetProps) {
  const navigate = useNavigate();
  function startTraining(id: number) {
    navigate(`/training/${id}`);
  }
  return (
    <Block>
      <h4 className="text-lg font-semibold">Trainings</h4>
      {!trainings || trainings.length === 0 ? (
        <p>No trainings</p>
      ) : (
        <div className="flex flex-col gap-4">
          {trainings!.map((v, i) => (
            <TrainingItem
              training={v}
              key={i}
              onClick={() => startTraining(v.id)}
            />
          ))}
        </div>
      )}
      <Button className="w-full bg-orange-200">Create basic training</Button>
    </Block>
  );
}
