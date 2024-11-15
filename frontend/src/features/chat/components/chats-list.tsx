import { CHAT_LIST } from "@/constants/cache";
import { useQuery } from "@tanstack/react-query";
import { getUserChatList } from "../services/chat.service";

export const ChatsList = () => {
  const { data, isLoading } = useQuery({
    queryKey: [CHAT_LIST],
    queryFn: getUserChatList,
  });
  return (
    <div className="mt-2">
      {!isLoading && data && (
        <ul>
          {data.map(({ id, messages }) => (
            <li key={id}>{messages[0]?.message_text || ""}</li>
          ))}
        </ul>
      )}
    </div>
  );
};
