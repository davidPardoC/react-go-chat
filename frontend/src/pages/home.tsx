import { ChatsList } from "@/features/chat/components/chats-list";
import { Input } from "@/components/ui/input";
import SearchUserModal from "@/features/chat/components/search-user-modal";

export const HomePage = () => {
  return (
    <div className="container p-2">
      <div className="flex gap-2 w-full">
        <Input placeholder="Search" className="w-full" />
        <SearchUserModal />
      </div>
      <ChatsList />
    </div>
  );
};
