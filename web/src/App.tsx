import { BrowserRouter as Router, Routes, Route } from "react-router-dom"

import AuthLayout from "./components/layouts/auth-layout";

import SignInPage from "./pages/signin";
import SignUpPage from "./pages/signup";

const App = () => {
  return (
    <Router>
      <Routes>
        <Route path="/auth" element={<AuthLayout />}>
          <Route path="signin" element={<SignInPage />} />
          <Route path="signup" element={<SignUpPage />} />
        </Route>
      </Routes>
    </Router>
  )
};

export default App;
