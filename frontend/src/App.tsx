import { Route, Switch } from "wouter";
import { HomePage } from "./pages/home";
import LoginPage from "./pages/login";
import SignupPage from "./pages/signup";
import { setAxiosDefaults } from "./utils/axios";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";

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
        <Route path="/">
          <HomePage />
        </Route>
      </Switch>
    </QueryClientProvider>
  );
}

export default App;
