import { Input } from "@/components/shared/Input";
import css from './CardItem.module.css';

export function CardItem({ position, question, answer, onUpdate }: {
    position: number,
    question: string,
    answer: string,
    onUpdate: (question: string, answer: string) => void
}) {
    return <div className={css.card}>
        <div className={css.position}>
            <a>#{position}</a>
        </div>
        <div className={css.content}>
            <Input placeholder="question" value={question} onInput={(e) => onUpdate(e.currentTarget.value, answer)} />
            <Input placeholder="answer" value={answer} onInput={(e) => onUpdate(question, e.currentTarget.value)} />
        </div>
    </div>
}