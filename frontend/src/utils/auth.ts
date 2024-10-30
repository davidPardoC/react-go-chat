import { Credentials } from "@/entities/credentials";

export const setCredentials = (credentials: Credentials) => {
  localStorage.setItem("acces_token", credentials.acces_token);
  localStorage.setItem("refresh_token", credentials.refresh_token);
};
