import { Credentials } from "@/entities/credentials";
import { jwtDecode } from "jwt-decode";

interface TokenPayload {
  email: string;
  exp: number;
  iss: string;
  sub: string;
  username: string;
}

export const setCredentials = (credentials: Credentials) => {
  localStorage.setItem("acces_token", credentials.acces_token);
  localStorage.setItem("refresh_token", credentials.refresh_token);
};

export const getCredentials = (): Credentials => {
  return {
    acces_token: localStorage.getItem("acces_token") || "",
    refresh_token: localStorage.getItem("refresh_token") || "",
  };
};

export const getTokenPayload = (): TokenPayload => {
  const token = getCredentials().acces_token;
  return jwtDecode(token);
};
