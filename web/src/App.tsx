import { BrowserRouter as Router, Routes, Route } from "react-router-dom"

import AuthLayout from "./components/layouts/auth-layout";
import GlobalLayout from "./components/layouts/global-layout";

import SignInPage from "./pages/signin";
import SignUpPage from "./pages/signup";
import MainPage from "./pages/main";

const App = () => {
  return (
    <Router>
      <Routes>
        <Route path="/auth" element={<AuthLayout />}>
          <Route path="signin" element={<SignInPage />} />
          <Route path="signup" element={<SignUpPage />} />
        </Route>
        <Route path="/" element={<GlobalLayout />}>
          <Route index element={<MainPage />} />
        </Route>

      </Routes>
    </Router>
  )
};

export default App;
