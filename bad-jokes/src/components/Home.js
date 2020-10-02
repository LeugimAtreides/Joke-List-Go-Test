import React from "react";

export default function Home() {
  const authenticate = () => null;

  return (
    <div>
      <div>
        <h1>Jadiel's Jokes</h1>
        <p>Sign in to get access</p>
        <a onClick={() => authenticate()}>Sign In</a>
      </div>
    </div>
  );
}
