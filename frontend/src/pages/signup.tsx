import { SignupForm } from "@/features/auth/signup/signup-form";

const SignupPage = () => {
  return (
    <main className="container mx-auto px-4 py-4">
      <h1 className="text-xl font-bold text-center">Go Chat App</h1>
      <SignupForm />
    </main>
  );
};

export default SignupPage;
