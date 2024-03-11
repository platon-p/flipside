import { CardSetRepository } from "@/repository/CardSetRepository";
import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { useAuth } from "@/hooks/Auth";
import Navbar from "../Navbar";
import { Input } from "@/components/shared/Input";
import { Button } from "@/components/shared/Button";
import './CreateSetPage.css';

export function CreateSetPage() {
    const navigate = useNavigate();
    const { userId, logout } = useAuth();
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

    return (<div><div className="header">
        <div className="logo"></div>
        <div className="sign-in">
            <a onClick={logout} href="/">выйти</a>
        </div>
    </div>

        <div className="naming">
            <h1 style={{marginBottom: 0}}>Создать набор</h1>
        </div>

        <div className="input-holder">
            <Input className='set-value' value={title} onInput={(e: any) => setTitle(e.currentTarget.value)} placeholder="название" />
            <div className="line"></div>
            <Input className='set-value' value={slug} onInput={(e: any) => setSlug(e.currentTarget.value)} placeholder="slug" />
            <div className="line"></div>
            <Button onClick={submit} className="create-set">продолжить</Button>
        </div>
        {errorMessage && <p 
        className="error" 
        style={{fontFamily: 'inter', fontSize: 12, marginTop:12 }}>
            {errorMessage}
            </p>}

        <Navbar />
    </div>

    )
}