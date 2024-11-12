import { ChatsList } from "@/features/chat/components/chats-list";
import { useWebSocket } from "@/features/chat/hooks/useWebsocket";
import { useEffect } from "react";

export const HomePage = () => {

  const socket = useWebSocket();

  useEffect(()=>{
    socket.onmessage = (event) => {
      console.log("Message received from server", event.data);
    };
  },[socket])

  socket.onopen = () => {
    console.log("Connected to websocket");
  };

  socket.onerror = (error) => {
    console.error("Websocket error", error);
  };

  return (
    <div className="flex p-5">
      <ChatsList />
    </div>
  );
};
