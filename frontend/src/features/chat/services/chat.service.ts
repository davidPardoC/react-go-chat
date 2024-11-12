import axios from "axios";
import { ChatList } from "../interfaces/chat-list.interface";

export const getUserChatList = async ():Promise<ChatList[]> => {
  const { data } = await axios.get("/v1/chats/");
  return data;
};
