import { Credentials } from "@/entities/credentials";

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
