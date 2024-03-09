import { CardSetRepository } from "@/repository/CardSetRepository";
import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { Input } from "@/components/shared/Input";
import { Button } from "@/components/shared/Button";
import './CreateSetPage.css';

export function CreateSetPage() {
    const navigate = useNavigate();
    const [title, setTitle] = useState('');
    const [slug, setSlug] = useState('');
    const [errorMessage, setErrorMessage] = useState('');

    function submit() {
        CardSetRepository.createCardSet(title, slug)
            .then((response) => {
                navigate(`/set/${response.slug}`)
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