import React from 'react';

function Profile() {
  return (
      <img
          className="w3-round"
          src="https://i.imgur.com/MK3eW3As.jpg"
          alt="Katherine Johnson"
      />
  );
}

export default function Gallery() {
  return (
      <section>
        <h1>Amazing scientists</h1>
        <Profile />
        <Profile />
        <Profile />
      </section>
  );
}