export const useWebSocket = () => {
  const socket = new WebSocket("ws://localhost:5000/ws");
  return socket;
};
