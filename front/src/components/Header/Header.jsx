"use client"

import { useState, useEffect } from "react"
import { MessageSquare } from "lucide-react"
import "./Header.css";


export default function Header() {
  const [scrolled, setScrolled] = useState(false)

  // Add scroll effect to change header appearance on scroll
  useEffect(() => {
    const handleScroll = () => {
      setScrolled(window.scrollY > 10)
    }

    window.addEventListener("scroll", handleScroll)
    return () => window.removeEventListener("scroll", handleScroll)
  }, [])

  return (
    <header className={`header ${scrolled ? "scrolled" : ""}`}>
      <div className="header-container">
        {/* Logo and Title */}
        <div className="logo-container">
          <div className="icon-wrapper">
            <MessageSquare className="chat-icon" />
            <span className="status-indicator"></span>
          </div>
          <h1 className="title">Besto Chat</h1>
        </div>

        {/* Actions */}
        <div className="actions-container">
          <button className="action-button primary">New Chat</button>
        </div>
      </div>
    </header>
  )
}

