import { Button } from "@/shared/Button";
import { TaskResult, TrainingTask } from "@/service/ApiService/ApiService";
import { TrainingService } from "@/service/TrainingService";
import { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";

export function TrainingPage() {
    const { id: trainingId } = useParams();
    const [task, setTask] = useState<TrainingTask | undefined>();
    const [result, setResult] = useState<TaskResult | undefined>();
    const navigate = useNavigate();

    function loadTask(trainingId: number) {
        TrainingService.getNextTask(trainingId)
            .then(task => {
                setTask(task);
            }).catch(e => {
                if (e === 'Training is completed') {
                    navigate(-1)
                    alert('Training is completed');
                } else {
                    console.error(e);
                }
            })
    }

    function submit(answer: 'Know' | "Don't know") {
        if (!task) {
            return;
        }
        TrainingService.submitAnswer(parseInt(trainingId!), answer)
            .then((res) => {
                setResult(res);
                loadTask(parseInt(trainingId!));
            }).catch(e => {
                console.error(e);
            })
    }

    useEffect(() => {
        Promise.resolve(loadTask(
            parseInt(trainingId!)
        ));
    }, [trainingId]);
    if (!task) {
        return <div>Loading...</div>
    }
    return <div style={{
        display: 'flex',
        flexDirection: 'column',
        padding: '2em'
    }}>
        <div style={{
            backgroundColor: 'lightgray',
            minHeight: '8em',
            width: '100%',
        }}>
            <p>{task?.question}</p>    
        </div>
        <br />
        <div style={{
            display: 'flex',
            flexDirection: 'row',
            justifyContent: 'space-between',
            gap: '1em',
        }}>
            <Button onClick={() => submit('Know')} style={{ width: '50%'}}>Know</Button>
            <Button onClick={() => submit("Don't know")} style={{ width: '50%'}}>Don't know</Button>
        </div>
        {result && <p>{result.is_correct ? 'Correct' : 'Incorrect'}</p>}
    </div>
}