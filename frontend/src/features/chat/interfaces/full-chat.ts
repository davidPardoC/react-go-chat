import { ChatMembers, Messages } from "./chat-list.interface";

export interface FullChat {
    id: number;
    messages: Messages[];
    chat_members: ChatMembers[];
}