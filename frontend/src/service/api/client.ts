import { AuthService } from "../AuthService";

import ky, { KyRequest, NormalizedOptions, KyResponse } from "ky";

export const config = {
  baseUrl: "http://localhost:80",
  auth: "api/auth",
  cardSet: "api/cardset",
  cards: "api/cards",
  training: "api/training",
};

async function after401(
  req: KyRequest,
  opt: NormalizedOptions,
  res: KyResponse,
): Promise<Response | undefined> {
  if (res.status !== 401) return res;
  await AuthService.loginByRefreshToken();
  const token = localStorage.getItem("accessToken");
  req.headers.set("Authorization", `Bearer ${token}`);
  return ky(req);
}

async function afterError(
  req: KyRequest,
  opt: NormalizedOptions,
  res: KyResponse,
): Promise<KyResponse> {
  if (res.ok) {
    return res;
  }

  const err = await res.json<MessageResponse>();
  throw new Error(err.message);
}

export const client = ky.create({
  prefixUrl: config.baseUrl,
  hooks: {
    afterResponse: [after401],
    // TODO: fix afterError hook
  },
});

export interface MessageResponse {
  status_code: number;
  message: string;
}
