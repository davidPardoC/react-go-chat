import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import UsersList from "@/features/users/components/users-list";
import { MessageSquarePlus } from "lucide-react";

const SearchUserModal = () => {
  return (
    <Dialog>
      <DialogTrigger asChild>
        <Button className="w-1/4">
          <MessageSquarePlus />
        </Button>
      </DialogTrigger>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>New Chat</DialogTitle>
            <div>
              <UsersList />
            </div>
        </DialogHeader>
      </DialogContent>
    </Dialog>
  );
};

export default SearchUserModal;
