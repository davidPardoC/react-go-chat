import { getCredentials } from "@/utils/auth";

export const useWebSocket = () => {
  const credentials = getCredentials();
  const socket = new WebSocket(
    `ws://localhost:5500/ws?token=${credentials.acces_token}`
  );

  return socket;
};
