import { CHAT_LIST } from "@/constants/cache";
import { useQuery } from "@tanstack/react-query";
import { getUserChatList } from "../services/chat.service";
import { Avatar, AvatarFallback } from "@/components/ui/avatar";
import { ChatMembers } from "../interfaces/chat-list.interface";
import { getTokenPayload } from "@/utils/auth";
import { Link } from "wouter";

const findTheOtherUser = (
  chat_members: ChatMembers[],
  currentUserId: number
) => {
  if (chat_members.length === 1) {
    return chat_members[0];
  }
  return chat_members.find((member) => member.user_id !== currentUserId);
};

export const ChatsList = () => {
  const { data, isLoading } = useQuery({
    queryKey: [CHAT_LIST],
    queryFn: getUserChatList,
  });

  const { sub } = getTokenPayload();
  const currentUserId = parseInt(sub);

  return (
    <div className="mt-2 ">
      {!isLoading && data && (
        <ul className="border-b-2">
          {data.map(({ id, messages, chat_members }) => (
            <Link href={`/chat/?recipient_id=${findTheOtherUser(chat_members, currentUserId)?.user_id}`} key={id}>
              <li key={id} className="border-t-2 p-2 flex items-center gap-2">
                <Avatar>
                  <AvatarFallback>
                    {findTheOtherUser(chat_members, parseInt(sub))
                      ?.user?.username.slice(0, 2)
                      .toUpperCase()}
                  </AvatarFallback>
                </Avatar>
                <div>
                  <h4 className="font-bold">
                    {findTheOtherUser(chat_members, currentUserId)?.user
                      .username || ""}
                  </h4>
                  <span className="text-sm text-gray-500">
                    {messages[0].message_text || ""}
                  </span>
                </div>
              </li>
            </Link>
          ))}
        </ul>
      )}
    </div>
  );
};
