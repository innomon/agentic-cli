# Claude Code (Go Port) — Master Architecture Overview

> **Primary Language:** Go
> **Framework:** `agentic` (Wrapper for Google ADK / GenAI)
> **UI Framework:** `tview` / `tcell` (TUI)
> **Runtime Target:** Native Binary

---

## 1. Vision
The Go port of Claude Code aims to provide a high-performance, native CLI assistant that replicates the agentic capabilities of the original TypeScript version while leveraging the `agentic` framework for robust tool integration and session management.

## 2. Core Subsystems

### 2.1 TUI Console (`pkg/console`)
- **Framework**: Built on `tview` for layout and `tcell` for terminal primitives.
- **Components**:
    - `ChatView`: Scrollable region for message history with Markdown support.
    - `InputField`: Multi-line capable input for user prompts and slash commands.
    - `StatusLine`: Dynamic footer showing model, token usage, and session state.
    - `ContextViz`: Visualization of context window usage.

### 2.2 Agentic Wrapper (`pkg/agent`)
- **Framework**: Integrates with `google.golang.org/adk` and `github.com/innomon/agentic`.
- **Registry**: Manages tool registration and agent loading.
- **Runner**: Orchestrates the interaction between the TUI, the LLM, and the tools.

### 2.3 Tool System (`pkg/tools`)
- **Interface**: Implements the `agent.Tool` interface from the ADK.
- **Permissions**: Layered permission checks (Automatic, Ask Once, Ask Always, Deny) managed via configuration.
- **Tools**: Includes `BashTool`, `FileReadTool`, `FileEditTool`, `GlobTool`, `WebFetchTool`, etc.

### 2.4 Command System (`pkg/commands`)
- **Slash Commands**: Handled by the TUI input processor.
- **Registry**: A mapping of command keywords (e.g., `/compact`) to Go handler functions.

### 2.5 Session & Context (`pkg/session`)
- **Persistence**: Conversation history stored in JSONL format.
- **Compaction**: Automatic context summarization when token limits are approached.
- **Memory**: Support for `CLAUDE.md` and project-specific memory files.

## 3. Data Flow

1. **User Input**: Captured in `InputField`.
2. **Command Dispatch**: If input starts with `/`, it's routed to the command registry.
3. **Agent Run**: Regular prompts are sent to `runner.Run`.
4. **Streaming**: Response events are streamed back to the `ChatView`.
5. **Tool Use**: If the model requests a tool, the `Runner` executes the corresponding Go tool (after optional user approval via TUI dialog).
6. **Persistence**: Every turn is recorded to the session transcript.

## 4. Configuration
- **Location**: `~/.agentic/config.yaml` or project-local `.agentic.yaml`.
- **Content**: Model selection, API keys (if not via OAuth), tool-specific settings, and TUI themes.
