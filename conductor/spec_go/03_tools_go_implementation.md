# Claude Code (Go Port) — Tool Implementations

This spec defines the porting of core Claude Code tools to the Go `agentic` framework.

## 1. Filesystem Tools

### `ReadTool`
- **Go Name**: `filesystem.Read`
- **Capability**: Reads files, handles large content by returning line-delimited slices.
- **Enhanced Features**: Detects MIME types (PDF, images) and returns appropriate Go structs for the runner to process.

### `WriteTool` / `EditTool`
- **Go Name**: `filesystem.Write`, `filesystem.Edit`
- **Logic**: Implements atomic writes and `sed`-like editing functionality.
- **Safety**: Integrates with the TUI to show a diff before applying changes.

## 2. Shell Execution Tools

### `BashTool`
- **Go Name**: `shell.Bash`
- **Logic**: Executes commands using `os/exec`.
- **Background Support**: Implements a task registry in Go to manage non-blocking commands.
- **Sandboxing**: Optionally uses `bwrap` (on Linux) or similar mechanisms to isolate execution.

## 3. Search Tools

### `GlobTool`
- **Go Name**: `search.Glob`
- **Logic**: Uses `path/filepath.Glob` or a more advanced Go glob library to list files.

### `GrepTool`
- **Go Name**: `search.Grep`
- **Logic**: Wraps the `rg` (ripgrep) binary if available, or implements a native Go grep.

## 4. Web Tools

### `WebFetchTool`
- **Go Name**: `web.Fetch`
- **Logic**: Uses `net/http` to retrieve content. Integrates a Markdown converter (e.g., `go-md2man` or custom) to provide clean text to the model.

### `WebSearchTool`
- **Go Name**: `web.Search`
- **Logic**: Integrates with Google Search API or Brave Search API.

## 5. Agent Tools

### `AgentTool` (Subagents)
- **Logic**: Allows the model to spawn a new `runner.Runner` with a specific prompt.
- **Isolation**: Subagents run in a separate Go routine with their own context and session history.
- **Messaging**: Implements a mailbox-style communication between parent and child agents.

## 6. Tool Permission Middleware
A Go middleware pattern will be used to wrap all tools for permission enforcement.

```go
func PermissionMiddleware(t agent.Tool) agent.Tool {
    // Check global config for auto-allow patterns
    // If not allowed, signal the TUI to show a confirmation dialog
    // Block execution until user responds
}
```
