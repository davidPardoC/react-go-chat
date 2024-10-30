import { Route, Switch } from "wouter";
import { HomePage } from "./pages/home";
import LoginPage from "./pages/login";
import SignupPage from "./pages/signup";
import { setAxiosDefaults } from "./utils/axios";

setAxiosDefaults();

function App() {
  return (
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
  );
}

export default App;
