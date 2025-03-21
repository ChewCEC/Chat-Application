"use client";

import React, { useState } from "react";
import useWebSocket from "./useWebSocket";

const ChatPage = () => {
  const { messages, sendMessage } = useWebSocket("ws://localhost:8080/ws");
  const [input, setInput] = useState("");

  const handleSend = () => {
    if (input.trim()) {
      sendMessage(input);
      setInput("");
    }
  };

  return (
    <div className="chat-container">
      <h2>Chat</h2>
      <div className="chat-messages">
        {messages.map((msg, index) => (
          <div key={index} className="message">
            {msg}
          </div>
        ))}
      </div>
      <input
        type="text"
        value={input}
        onChange={(e) => setInput(e.target.value)}
        placeholder="Type a message..."
      />
      <button onClick={handleSend}>Send</button>
    </div>
  );
};

export default ChatPage;
