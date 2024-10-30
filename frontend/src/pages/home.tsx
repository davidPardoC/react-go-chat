import { Button } from "@/components/ui/button";
import { useWebSocket } from "@/features/chat/hooks/useWebsocket";

export const HomePage = () => {
  const socket = useWebSocket();

  socket.onopen = () => {
    console.log("Connected to websocket");
  };

  return (
    <div>
      <Button
        onClick={() => {
          socket.send("Hello from the frontend!");
        }}
      >
        Send
      </Button>
    </div>
  );
};
