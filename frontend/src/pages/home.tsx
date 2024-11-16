import { ChatsList } from "@/features/chat/components/chats-list";
import { Input } from "@/components/ui/input";
import SearchUserModal from "@/features/chat/components/search-user-modal";

export const HomePage = () => {
  return (
    <div className="container">
      <div className="flex gap-2 w-full p-3">
        <Input placeholder="Search" className="w-full" />
        <SearchUserModal />
      </div>
      <ChatsList />
    </div>
  );
};
