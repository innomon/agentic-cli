# Release Notes - v0.0.0 (Initial Release)

We are excited to announce the initial release of **AgentiCli**, a high-performance, native Go terminal assistant built on the `agentic` framework (Google ADK). AgentiCli brings powerful agentic capabilities to your terminal with a rich, interactive TUI.

## ✨ Key Features

### 🖥️ Rich Terminal User Interface (TUI)
- Built with `tview` and `tcell` for a responsive, multi-region layout.
- Real-time streaming of agent responses and tool execution steps.
- Interactive components for input, chat history, and status monitoring.

### 🤖 Agentic Core
- Deep integration with the **Google ADK (Agentic Development Kit)**.
- Robust session management and tool execution workflows.
- Native performance with a single Go binary.

### 🛠️ Built-in Local Tools
- **Filesystem**: Safe reading, writing, and surgical editing of files with ambiguity detection.
- **Shell**: Execute bash or CMD commands directly from the agent.
- **Search**: Powerful recursive globbing and regex-based grep capabilities.
- **Web**: Fetch and process remote content via HTTP.

### ⌨️ Smart Interaction & Shortcuts
- **Prompt History**: Cycle through previous prompts with `Up`/`Down` arrows.
- **History Picker**: Fuzzy search through your session history with `Ctrl+R`.
- **File Attachments**: Use the `@/path/to/file` syntax to attach text, PDFs, or images to your prompts.
- **Slash Commands**: Quick access to utility functions like `/help`, `/clear`, `/history`, and `/exit`.

## 📦 Getting Started

### Prerequisites
- Go 1.25.6 or higher.

### Installation
```bash
go build -o agenticli main.go
./agenticli
```

## 📝 License
This project is licensed under the Apache License, Version 2.0.

---
*Built with passion using the Agentic Framework.*
