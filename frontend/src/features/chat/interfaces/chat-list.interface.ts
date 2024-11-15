export interface ChatList {
  id: number;
  messages: Messages[];
}

export interface Messages {
  id: number;
  message_text: string;
  read: boolean;
  chat_id: number;
  user_id: number;
  created_at: string;
  updated_at: string;
}
