import { useState, useEffect, useRef } from "react";

const useWebSocket = (url: string) => {
  const [messages, setMessages] = useState<string[]>([]);
  const socketRef = useRef<WebSocket | null>(null);

  useEffect(() => {
    const socket: WebSocket = new WebSocket(url);
    socketRef.current = socket;

    socket.onopen = () => {
      console.log("Connected to WebSocket");
    };

    socket.onmessage = (event) => {
      setMessages((prev) => [...prev, event.data]);
    };

    socket.onclose = () => {
      console.log("WebSocket disconnected");
    };

    socket.onerror = (error) => {
      console.log("WebSocket error:", error);
    };

    return () => {
      socket.close();
    };
  }, [url]);

  const sendMessage = (message: string) => {
    if (socketRef.current && socketRef.current.readyState === WebSocket.OPEN) {
      socketRef.current.send(message);
    }
  };

  return { messages, sendMessage };
};

export default useWebSocket;
