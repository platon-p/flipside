import { ApiService } from "@/service/ApiService";
import { useMemo } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { create, useStore } from "zustand";
import { CardSetList } from "../MainPage/UserPage/CardSetList";
import { CardSet } from "@/repository/CardSetRepository";

interface ProfileStore {
  profile?: {
    name: string;
    nickname: string;
  };
  cardsets: CardSet[];
  loadProfile(nickname: string): Promise<any>;
  loadCardsets(nickname: string): Promise<any>;
}
const st = create<ProfileStore>((set, get) => {
  return {
    profile: undefined,
    cardsets: [],

    async loadProfile(nickname: string) {
      set({
        profile: await ApiService.Profile.getUserProfile(nickname),
      });
    },
    async loadCardsets(nickname: string) {
      set({
        cardsets: await ApiService.Profile.getUserCardSets(nickname),
      });
    },
  };
});

export default function Profile() {
  const { nickname } = useParams();
  const { profile, loadProfile } = useStore(st);
  useMemo(() => {
    loadProfile(nickname!);
  }, [nickname]);
  if (!profile) return;
  return (
    <div className="mt-8 mx-auto w-1/2 flex flex-col">
      <h1 className="text-2xl font-bold">{profile.name}</h1>
      <p>@{profile.nickname}</p>
      <div className="h-4"></div>
      <h2 className="text-xl font-medium">Наборы</h2>
      <CardSets />
    </div>
  );
}

function CardSets() {
  const { nickname } = useParams();
  const navigate = useNavigate();
  const { cardsets, loadCardsets } = useStore(st);
  useMemo(() => {
    loadCardsets(nickname!);
  }, [nickname]);
  function openCardset(slug: string) {
    navigate(`/set/${slug}`);
  }
  return <CardSetList cards={cardsets} onClick={openCardset} />;
}
