import { BrowserRouter as Router, Routes, Route } from "react-router-dom"

import AuthLayout from "./components/layouts/auth-layout";

import SignInPage from "./pages/signin";

const App = () => {
  return (
    <Router>
      <Routes>
        <Route path="/auth" element={<AuthLayout />}>
          <Route path="signin" element={<SignInPage />} />
        </Route>
      </Routes>
    </Router>
  )
};

export default App;
