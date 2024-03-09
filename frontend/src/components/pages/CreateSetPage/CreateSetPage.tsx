import { Input } from "@/components/shared/Input";
import { useState } from "react";
import './CreateSetPage.css';
import { Button } from "@/components/shared/Button";
import { CardSetRepository } from "@/repository/CardSetRepository";

export function CreateSetPage() {
    const [title, setTitle] = useState('');
    const [slug, setSlug] = useState('');
    const [errorMessage, setErrorMessage] = useState('');

    function submit() {
        CardSetRepository.createCardSet(title, slug)
            .then((response) => {
                console.log(response);
            })
            .catch(err => {
                setErrorMessage(err?.toString());
            })
    }

    return (
        <div>
            <h1>Create a Set</h1>
            <div className="input-holder">
                <Input value={title} onInput={e => setTitle(e.currentTarget.value)} placeholder="title" />
                <Input value={slug} onInput={e => setSlug(e.currentTarget.value)} placeholder="slug" />
                <Button onClick={submit}>Continue</Button>
            </div>
            {errorMessage && <p className="error">{errorMessage}</p>}
        </div>
    )
}