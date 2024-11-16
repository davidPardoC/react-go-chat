import axios from "axios";
import { ChatList } from "../interfaces/chat-list.interface";
import { FullChat } from "../interfaces/full-chat";

export const getUserChatList = async ():Promise<ChatList[]> => {
  const { data } = await axios.get("/v1/chats/");
  return data;
};

export const getFullChat = async (chatId: number):Promise<FullChat> => {
  const { data } = await axios.get(`/v1/chats/${chatId}`);
  return data;
}