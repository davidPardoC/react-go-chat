import { Avatar, AvatarFallback } from "@/components/ui/avatar";
import { Button } from "@/components/ui/button";
import { Form, FormField } from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { ChatMessage } from "@/features/chat/interfaces/message.interface";
import { sendTextMessage } from "@/features/chat/services/realtime.chat.service";
import { getCredentials } from "@/utils/auth";
import { zodResolver } from "@hookform/resolvers/zod";
import { SendHorizonal } from "lucide-react";
import { useEffect, useRef, useState } from "react";
import { useForm } from "react-hook-form";
import { Redirect } from "wouter";
import { z } from "zod";

const formSchema = z.object({
  message: z.string(),
});

type FormValues = z.infer<typeof formSchema>;

const ChatPage = () => {
  const urlParams = new URLSearchParams(window.location.search);
  const recipientId = urlParams.get("recipient_id");
  const [messages, setMessages] = useState<ChatMessage[]>([]);
  const socketRef = useRef<WebSocket | null>(null);

  const form = useForm<FormValues>({
    resolver: zodResolver(formSchema),
    defaultValues: { message: "" },
  });

  useEffect(() => {
    const socket = new WebSocket(
      `ws://localhost:5500/ws?token=${getCredentials().acces_token}`
    );

    socket.addEventListener("message", (event) => {
      setMessages((prev) => [...prev, JSON.parse(event.data)]);
    });

    socketRef.current = socket;

    return () => socket.close();
  }, []);

  useEffect(() => {
    console.log(messages);
  }, [messages]);

  if (!recipientId) {
    return <Redirect to="/" />;
  }

  const onSubmit = (values: FormValues) => {
    if (socketRef.current) {
      sendTextMessage(socketRef.current, parseInt(recipientId), values.message);
      form.setValue("message", "");
    }
  };

  return (
    <div className="container">
      <div className="h-screen">
        <div className="px-2">
          {messages.map((message, index) => (
            <div key={index} className="flex items-center gap-1 mt-4">
              <Avatar>
                <AvatarFallback>
                  S
                </AvatarFallback>
              </Avatar>
              <span className="p-1 bg-gray-200 w-max px-3 rounded-xl">
                {message.data.message_text}
              </span>
            </div>
          ))}
        </div>
      </div>
      <Form {...form}>
        <form
          className="flex gap-2 absolute bottom-0 pb-3 w-full mx-auto px-2"
          onSubmit={form.handleSubmit(onSubmit)}
        >
          <FormField
            control={form.control}
            name="message"
            render={({ field }) => (
              <Input
                {...field}
                placeholder="Some text..."
                className="rounded-full"
              />
            )}
          />
          <Button type="submit" className="rounded-full">
            <SendHorizonal />
          </Button>
        </form>
      </Form>
    </div>
  );
};

export default ChatPage;
