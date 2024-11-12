import { CHAT_LIST } from "@/constants/cache";
import { useQuery } from "@tanstack/react-query";
import { getUserChatList } from "../services/chat.service";

export const ChatsList = () => {
  const { data, isLoading } = useQuery({
    queryKey: [CHAT_LIST],
    queryFn: getUserChatList,
  });
  return (
    <>
      {!isLoading && data && (
        <ul>
          {data.map(({id}) => (
            <li key={id}></li>
          ))}
        </ul>
      )}
    </>
  );
};
