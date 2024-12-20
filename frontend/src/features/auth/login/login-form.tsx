import { Alert, AlertDescription, AlertTitle } from "@/components/ui/alert";
import { Button } from "@/components/ui/button";
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { setCredentials } from "@/utils/auth";
import { zodResolver } from "@hookform/resolvers/zod";
import { AxiosError } from "axios";
import { useState } from "react";
import { useForm } from "react-hook-form";
import { useLocation } from "wouter";
import { loginUser } from "../services/auth.services";
import { setAxiosDefaults } from "@/utils/axios";
import { z } from "zod";

const formSchema = z.object({
  email: z.string().email(),
  password: z.string().min(8, {
    message: "Password must be at least 8 characters.",
  }),
});

export type LoginForm = z.infer<typeof formSchema>;

export const LoginForm = () => {
  const form = useForm<LoginForm>({ resolver: zodResolver(formSchema) });
  const [, setLocation] = useLocation();
  const [error, setError] = useState<string>();

  async function onSubmit(values: LoginForm) {
    try {
      const credentials = await loginUser(values);
      setCredentials(credentials);
      setAxiosDefaults()
      setLocation("/");
    } catch (error:unknown) {
      const { response } = error as AxiosError;
      if (response && response.data) {
        const data = response.data as { error: string };
        setError(data.error);
      } else {
        setError("An unknown error occurred.");
      }
    }
  }

  return (
    <Form {...form}>
      {error && (
        <Alert variant="destructive">
          <AlertTitle>Error</AlertTitle>
          <AlertDescription>{error}</AlertDescription>
        </Alert>
      )}
      <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
        <FormField
          control={form.control}
          name="email"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Email</FormLabel>
              <FormControl>
                <Input placeholder="example@go.com" {...field} />
              </FormControl>
              <FormDescription>
                This is your public display name.
              </FormDescription>
              <FormMessage />
            </FormItem>
          )}
        />{" "}
        <FormField
          control={form.control}
          name="password"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Password</FormLabel>
              <FormControl>
                <Input type="password" {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
        <Button type="submit">Submit</Button>
      </form>
    </Form>
  );
};
