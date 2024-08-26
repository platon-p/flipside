import { AuthApi } from "./api/AuthApi";
import { CardApi } from "./api/CardApi";
import { CardSetApi } from "./api/CardSetApi";
import { ProfileApi } from "./api/ProfileApi";
import { TrainingApi } from "./api/TrainingApi";

export const ApiService = {
  Auth: AuthApi,
  CardSet: CardSetApi,
  Card: CardApi,
  Training: TrainingApi,
  Profile: ProfileApi,
};
