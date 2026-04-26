# AgentiCli (Go Port)

A dedicated TUI (Terminal User Interface) console for the `agentic` framework, ported from Claude Code specifications to a high-performance native Go implementation.

## Features
- **TUI Interface**: Rich, multi-region terminal UI powered by `tview` and `tcell`.
- **Local Tool Support**: Integrated filesystem (Read/Write/Edit), Shell (Bash/CMD), Search (Glob/Grep), and Web (Fetch) tools.
- **File Attachments**: Use `@/path/to/file` syntax to attach PDFs, images, and text files.
- **Slash Commands**: Specialized commands like `/help`, `/clear`, `/history`, and `/exit`.
- **Input History**: Navigation with arrows and a searchable history picker via `Ctrl+R`.
- **Streaming Support**: Real-time response and tool execution feedback.

## Project Structure
- `main.go`: Entry point, initializes the registry and launches the TUI console.
- `pkg/agent/`: Agent runner and session management.
- `pkg/commands/`: Slash command implementation and registry.
- `pkg/console/`: TUI layout and interactive components.
- `pkg/tools/`: ADK-compliant tool implementations.

## Development Workflows

### Build
```bash
go build -o agenticli main.go
```

### Run
```bash
./agenticli [config.yaml]
```
If no config is provided, default settings are used. See `config.yaml.example` for a template.

### Usage
- **Type message**: Type your message and press `Enter`.
- **Execute command**: Type `/` followed by the command name (e.g., `/help`).
- **Attach file**: Add `@path/to/file` anywhere in your message.
- **Navigate history**: Use `Up`/`Down` arrows or `Ctrl+R`.

### Coding Standards
- Use `go fmt` and `go mod tidy` for all changes.
- Add unit tests for new tools and core logic in `pkg/tools/` or `pkg/commands/`.
