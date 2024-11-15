export interface ChatMessage {
  data: { message_text: string; recipient_id: number; chat_id: number };
  event: string;
  user_id: number;
}
