import { Button } from "@/shared";
import { TaskResult, TrainingTask } from "@/service/ApiService/TrainingApi";
import { TrainingService } from "@/service/TrainingService";
import { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";

export function TrainingPage() {
  const navigate = useNavigate();
  const { id: trainingId } = useParams();
  const [task, setTask] = useState<TrainingTask | undefined>();
  const [result, setResult] = useState<TaskResult | undefined>();

  function loadTask(trainingId: number) {
    TrainingService.getNextTask(trainingId)
      .then((task) => {
        setTask(task);
      })
      .catch((e) => {
        if (e === "Training is completed") {
          navigate(-1);
          alert("Training is completed");
        } else {
          console.error(e);
        }
      });
  }

  function submit(answer: "Know" | "Don't know") {
    if (!task) {
      return;
    }
    TrainingService.submitAnswer(parseInt(trainingId!), answer)
      .then((res) => {
        setResult(res);
        loadTask(parseInt(trainingId!));
      })
      .catch((e) => {
        console.error(e);
      });
  }

  useEffect(() => {
    Promise.resolve(loadTask(parseInt(trainingId!)));
  }, [trainingId]);
  if (!task) {
    return <div>Loading...</div>;
  }
  return (
    <div className="flex flex-col p-2">
      <div
        className="w-full"
        style={{
          backgroundColor: "lightgray",
          minHeight: "8em",
        }}
      >
        <p>{task?.question}</p>
      </div>
      <br />
      <div className="flex justify-between gap-4">
        <Button onClick={() => submit("Know")} className="w-1/2">
          Know
        </Button>
        <Button onClick={() => submit("Don't know")} className="w-1/2">
          Don't know
        </Button>
      </div>
      {result && <p>{result.is_correct ? "Correct" : "Incorrect"}</p>}
    </div>
  );
}
