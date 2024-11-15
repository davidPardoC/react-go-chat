import { WebSocketEvents } from "@/constants/webscoket-events";
import { getTokenPayload } from "@/utils/auth";

export const sendTextMessage = (
  socket: WebSocket,
  recipientId: number,
  message: string,
) => {
  const { sub } = getTokenPayload();
  const payload = {
    event: WebSocketEvents.MESSAGE,
    user_id: parseInt(sub),
    data: {
      message_text: message,
      recipient_id: recipientId,
    },
  };
  socket.send(JSON.stringify(payload));
};
