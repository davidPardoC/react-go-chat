import { Button } from "@/components/ui/button";
import { useWebSocket } from "@/features/chat/hooks/useWebsocket";
import UsersList from "@/features/users/components/users-list";
import { getAllUsers } from "@/features/users/services/users.service";
import { useEffect } from "react";

export const HomePage = () => {

  const socket = useWebSocket();

  useEffect(()=>{
    socket.onmessage = (event) => {
      console.log("Message received from server", event.data);
    };
    return () => {
      socket.close()
    }
  },[])

  useEffect(()=>{
    getAllUsers();
  }, [])

  socket.onopen = () => {
    console.log("Connected to websocket");
  };

  socket.onerror = (error) => {
    console.error("Websocket error", error);
  };

  return (
    <div className="flex p-5">
      <UsersList />
      <Button
      className="ml-4"
        onClick={() => {
          socket.send("Hello from the frontend!");
        }}
      >
        Send
      </Button>
    </div>
  );
};
