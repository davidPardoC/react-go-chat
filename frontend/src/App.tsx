import { Route, Switch } from "wouter";
import { HomePage } from "./pages/home";
import LoginPage from "./pages/login";
import SignupPage from "./pages/signup";
import { setAxiosDefaults } from "./utils/axios";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import ProtectedRoute from "./components/ui/protected-route";
import ChatPage from "./pages/chat";

setAxiosDefaults();

const queryClient = new QueryClient()

function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <Switch>
        <Route path="/signup">
          <SignupPage />
        </Route>
        <Route path="/login">
          <LoginPage />
        </Route>
        <ProtectedRoute path="/">
          <HomePage />
        </ProtectedRoute>
        <ProtectedRoute path="/chat">
          <ChatPage />
        </ProtectedRoute>
      </Switch>
    </QueryClientProvider>
  );
}

export default App;
