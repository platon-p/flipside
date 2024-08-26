import { config, client } from "./client";

export interface RegisterRequest {
  name: string;
  nickname: string;
  email: string;
  password: string;
}

export interface TokenPairResponse {
  access_token: string;
  refresh_token: string;
  expires_in: Date;
}

export const AuthApi = {
  async register(request: RegisterRequest): Promise<TokenPairResponse> {
    const response = await client.post(`${config.auth}/register`, {
      json: request,
    });
    return await response.json();
  },

  async loginByEmail(
    email: string,
    password: string,
  ): Promise<TokenPairResponse> {
    const response = await client.post(`${config.auth}/login-by-email`, {
      json: { email, password },
    });
    return await response.json();
  },

  async loginByToken(token: string): Promise<TokenPairResponse> {
    const response = await client.post(`${config.auth}/login-by-token`, {
      json: { refresh_token: token },
    });
    return response.json();
  },
};
