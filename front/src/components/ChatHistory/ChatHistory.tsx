import React from "react";
import "./ChatHistory.css";

interface ChatHistoryProps {
  chatHistory: { data: string }[];
}

const ChatHistory: React.FC<ChatHistoryProps> = ({ chatHistory }) => {
  return (
    <div className="ChatHistory">
      <h2>Chat History</h2>
      {chatHistory.map((msg, index) => (
        <p key={index}>{msg.data}</p>
      ))}
    </div>
  );
};

export default ChatHistory;