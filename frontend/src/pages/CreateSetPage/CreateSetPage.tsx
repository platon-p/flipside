import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { CardSetRepository } from "@/repository/CardSetRepository";
import { useAuth } from "@/hooks/Auth";
import { Input, Button } from "@/shared";

export function CreateSetPage() {
  const navigate = useNavigate();
  const { logout } = useAuth();
  const [title, setTitle] = useState("");
  const [slug, setSlug] = useState("");
  const [errorMessage, setErrorMessage] = useState("");

  function submit() {
    CardSetRepository.createCardSet(title, slug)
      .then((response) => {
        navigate(`/set/${response.slug}`);
      })
      .catch((err) => {
        setErrorMessage(err?.toString());
      });
  }

  return (
    <div>
      <div className="flex flex-col gap-1 px-1">
        <a onClick={() => navigate(-1)}>Назад</a>
        <h1 className="mb-0 text-2xl font-medium">Создать набор</h1>
        <Input
          value={title}
          onInput={(e) => setTitle(e.currentTarget.value)}
          placeholder="название"
        />
        <div className="line"></div>
        <Input
          value={slug}
          onInput={(e) => setSlug(e.currentTarget.value)}
          placeholder="slug"
        />
        <div className="line"></div>
        <Button onClick={submit} className="create-set">
          продолжить
        </Button>
      </div>
      {errorMessage && <p className="text-red-500 text-lg">{errorMessage}</p>}
    </div>
  );
}
