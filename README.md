# AgentiCli (Go Port) - Agentic CLI

A high-performance, native Go terminal assistant built on the `agentic` framework (Google ADK). This tool replicates the powerful agentic capabilities of Claude Code in a native binary with a rich TUI.

## 🚀 Features

- **Rich TUI**: Powered by `tview` and `tcell`, providing a multi-region terminal interface.
- **Agentic Core**: Full integration with the `agentic` framework for robust tool execution and session management.
- **Local Tools**: Built-in support for:
    - **Filesystem**: Read, write, and safely edit files (detects ambiguous replacements).
    - **Shell**: Execute bash/cmd commands with standard output/error capture.
    - **Search**: Recursive globbing and regex-based grep.
    - **Web**: Fetch remote content via HTTP.
- **Slash Commands**:
    - `/help`: List all available commands.
    - `/clear`: Clear the chat history view.
    - `/history`: Access session history.
    - `/exit` or `/quit`: Exit the application.
- **Smart Input**:
    - **History Navigation**: Use `Up`/`Down` arrows to cycle through previous prompts.
    - **History Picker**: Press `Ctrl+R` to search and select from prompt history.
    - **File Attachments**: Use `@/path/to/file` syntax to attach images, PDFs, or text files.
- **Real-time Streaming**: Watch the agent think and execute tools in real-time.

## 🛠 Project Structure

- `main.go`: Application entry point and registry initialization.
- `pkg/agent/`: Wrapper for the `agentic` runner and session management.
- `pkg/commands/`: Implementation of slash commands.
- `pkg/console/`: Main TUI layout and interactive components.
- `pkg/tools/`: Local tool implementations (FS, Shell, Search, Web).
- `conductor/spec_go/`: Adapted architecture specifications.

## 📦 Getting Started

### Prerequisites

- Go 1.25.6 or higher.

### Installation & Build

```bash
# Clone the repository
git clone <repository-url>
cd agentic-cli

# Build the native binary
go build -o agenticli main.go
```

### Running

```bash
./agenticli [config.yaml]
```

## ⚙️ Configuration

AgentiCli uses a YAML configuration file to define models, agents, and tools. By default, it looks for `config.yaml` in the current directory if passed as an argument, or uses default settings.

To get started:

1.  **Copy the example configuration:**
    ```bash
    cp config.yaml.example config.yaml
    ```
2.  **Edit `config.yaml`:** Add your API keys (e.g., `api_key` for Gemini or OpenAI) and customize your default model or agent instructions.
3.  **Run with the config:**
    ```bash
    ./agenticli config.yaml
    ```

### Key Configuration Sections:

- **`models`**: Define LLM providers (Gemini, OpenAI, Ollama).
- **`agents`**: Configure agent behavior, instructions, and available tools.
- **`tools`**: Register local tools (filesystem, shell, search, web).
- **`root_agent`**: Specify which agent should handle the main conversation.

## ⌨️ Shortcuts

| Shortcut | Action |
|---|---|
| `Enter` | Submit your prompt or select a history item |
| `Up / Down` | Navigate prompt history |
| `Ctrl+R` | Open the fuzzy history search picker |
| `Esc` | Close the history picker or clear current input |
| `/exit` | Quit the application |

## 📝 Development

### Adding New Tools
1. Define your tool in `pkg/tools/`.
2. Implement the `New<Name>Tool()` function using `functiontool.New`.
3. Add it to the list in `pkg/tools/tools.go`.

### Running Tests
```bash
go test ./pkg/commands/... ./pkg/tools/...
```

---
*Built with the Agentic Framework.*

## 📄 License

This project is licensed under the Apache License, Version 2.0. See the [LICENSE](LICENSE) file for details.
