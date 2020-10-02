import React, { useState } from "react";

export default function Joke({ joke }) {
  const [liked, setLiked] = useState("");
  const like = () => null;
  return (
    <div>
      <div>
        <div>
          #{joke.id}
          <span>{liked}</span>
        </div>
        <div>
          {joke.likes} Likes &nbsp;
          <a onClick={() => like()}>
            <span>Thumbs Up!</span>
          </a>
        </div>
      </div>
    </div>
  );
}
