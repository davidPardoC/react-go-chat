export interface ChatList {
  id: number;
  messages: Messages[];
  chat_members: ChatMembers[];
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

export interface ChatMembers {
  id: number;
  user_id: number;
  chat_id: number;
  created_at: string;
  updated_at: string;
  user: User;
}

interface User {
  id: number;
  username: string;
}
