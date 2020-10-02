import React from "react";
import Home from "./components/Home";
import LoggedIn from "./components/LoggedIn";
import "./App.css";

function App() {
  const [loggedIn, setLoggedIn] = React.useState(true);
  if (loggedIn) return <LoggedIn />;
  else return <Home />;
}

export default App;
