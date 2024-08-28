import { Card } from "@/repository/CardRepository";
import { ApiService } from "@/service/ApiService";
import { create } from "zustand";

interface CardSetState {
  title: string;
  slug: string;
  cards: Card[];
  state: "idle" | "loading" | "error";
  error: string;

  fetchSet(slug: string): Promise<void>;
  fetchCards(slug: string): Promise<void>;

  setTitle(title: string): void;
  setSlug(slug: string): void;

  selectCard(position: number): Card;
  editCard(card: Card): void;

  addCard(): number;
}

export const useCardSet = create<CardSetState>((set, get) => ({
  title: "",
  slug: "",
  cards: [],
  state: "idle",
  error: "",
  fetchSet: (slug: string) => {
    return ApiService.CardSet.getCardSet(slug)
      .then((cardSet) => {
        set({ title: cardSet.title, slug: cardSet.slug });
      })
      .catch((error) => {
        set({ error: error.message, state: "error" });
      });
  },
  fetchCards: (slug: string) => {
    return ApiService.Card.getCards(slug)
      .then((cards) => {
        set({ cards });
      })
      .catch((error) => {
        set({ error: error.message, state: "error" });
      });
  },

  setTitle: (title: string) => set({ title }),
  setSlug: (slug: string) => set({ slug }),

  selectCard: (position: number) => {
    return get().cards.filter((card) => card.position === position)[0];
  },
  editCard(card: Card) {
    set((state) => {
      const i = state.cards.findIndex((v) => v.position == card.position);
      state.cards[i] = card;
      return { cards: state.cards };
    });
  },
  addCard(): number {
    const positions = get().cards.map((i) => i.position);
    const nextPosition = Math.max(...positions) + 1;
    set((state) => {
      state.cards.push({ position: nextPosition, question: "", answer: "" });
      return { cards: state.cards };
    });
    return nextPosition;
  },
}));
