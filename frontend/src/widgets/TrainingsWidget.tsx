import { useNavigate, useParams } from "react-router-dom";
import { TrainingRepository, TrainingSummary } from "@/repository/TrainingRepository";
import { Block, Button } from "@/shared";
import { TrainingItem } from "./TrainingItem";

interface TrainingsWidgetProps {
  trainings: TrainingSummary[] | undefined;
}

export function TrainingsWidget({ trainings }: TrainingsWidgetProps) {
  const navigate = useNavigate();
  const slug = useParams()['slug']!;

  function startTraining(id: number) {
    navigate(`/training/${id}`);
  }
  function createTraining(trainingType: string) {
    TrainingRepository.createTraining(slug, trainingType)
      .then((training) => {
        console.log("Training created", training);
        window.location.reload();
      })
      .catch((e) => {
        console.log(e);
      });
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
      <Button onClick={() => createTraining('basic')} className="w-full bg-orange-200">Create basic training</Button>
    </Block>
  );
}
