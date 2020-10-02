import React from "react";
import Joke from './Joke';

export default function LoggedIn() {
  const [jokes, setJokes] = React.useState([]);
  const logout = () => null;
  return (
    <div>
      <div>
        <br />
        <span>
          <a onClick={() => logout}>Logout</a>
        </span>
        <h2>Jadiel's Jokes</h2>
        <p>Vote on each joke Jadiel has made!</p>
        {jokes.map((joke) => (
          <Joke key={joke.id} joke={joke} />
        ))}
      </div>
    </div>
  );
}
