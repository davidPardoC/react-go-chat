import { LoginForm } from "@/features/auth/login/login-form";

const LoginPage = () => {
  return (
    <main className="container mx-auto px-4 py-4">
      <h1 className="text-xl font-bold text-center">Go Chat App</h1>
      <LoginForm />
    </main>
  );
};

export default LoginPage;
