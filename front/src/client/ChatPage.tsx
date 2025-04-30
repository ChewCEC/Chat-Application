"use client"

import type React from "react"
import { useState, useRef, useEffect } from "react"
import useWebSocket from "./useWebSocket"
import "./ChatPage.css"

interface Message {
  id: string;
  content: string;
  timestamp: string;
}

const ChatPage = () => {
  const { messages, sendMessage } = useWebSocket("ws://localhost:8080/ws")
  const [input, setInput] = useState("")
  const messagesEndRef = useRef<HTMLDivElement>(null)

  // Parse messages from JSON string to object
  const parsedMessages = messages.map(msg => {
    try {
      return JSON.parse(msg) as Message;
    } catch (e) {
      return { content: msg } as Message;
    }
  });

  // Auto-scroll to bottom when new messages arrive
  useEffect(() => {
    messagesEndRef.current?.scrollIntoView({ behavior: "smooth" })
  }, [messages])

  const handleSend = () => {
    if (input.trim()) {
      sendMessage(input)
      setInput("")
    }
  }

  // Handle Enter key press
  const handleKeyPress = (e: React.KeyboardEvent) => {
    if (e.key === "Enter") {
      handleSend()
    }
  }

  return (
    <div className="chat-container">
      <div className="chat-header">
        <h2>Besto Chat</h2>
        <span className="online-indicator">Online</span>
      </div>

      <div className="chat-messages">
        {parsedMessages.map((msg, index) => (
          <div key={msg.id || index} className={`message ${index % 2 === 0 ? "sent" : "sent"}`}>
            <div className="message-content">{msg.content}</div>
            <div className="message-time">
              {new Date(msg.timestamp || Date.now()).toLocaleTimeString([], { 
                hour: "2-digit", 
                minute: "2-digit" 
              })}
            </div>
          </div>
        ))}
        <div ref={messagesEndRef} />
      </div>

      <div className="chat-input-container">
        <input
          type="text"
          value={input}
          onChange={(e) => setInput(e.target.value)}
          onKeyUp={handleKeyPress}
          placeholder="Type a message..."
          className="chat-input"
        />
        <button onClick={handleSend} className="send-button" disabled={!input.trim()}>
          Send
        </button>
      </div>
    </div>
  )
}

export default ChatPage

