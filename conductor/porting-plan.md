# Implementation Plan - Porting Claude Code to Go (Agentic Framework)

This plan outlines the steps to adapt the Claude Code specifications into the current Go architecture using the `agentic` framework.

## 1. Objectives
- Port the rich feature set of Claude Code to a Go-based TUI application.
- Leverage the `agentic` framework (Google ADK wrapper) for agent logic, tool execution, and session management.
- Replace React/Ink with `tview`/`tcell` for the terminal interface.
- Maintain compatibility with Claude Code's core concepts: tools, slash commands, multi-agent delegation, and context management.

## 2. Architecture Alignment

| Feature | Claude Code (TS/Ink) | Go Architecture (Agentic) |
|---|---|---|
| **Language** | TypeScript / Bun | Go |
| **TUI** | React / Ink / Yoga | `tview` / `tcell` |
| **Agent Core** | `QueryEngine.ts` | `runner.New` / `agent.Run` |
| **Tools** | `Tool.ts` (Class-based) | `agent.Tool` (Interface-based) |
| **Commands** | Slash commands in `commands/` | Integrated into TUI input handler |
| **State** | React Context / AppState | `session.Session` / Go structs |
| **Config** | JSON/YAML in `~/.claude` | `config.Config` / `~/.agentic` |

## 3. Implementation Phases

### Phase 1: Core Foundation & adapted Specs
- [x] Create `spec_go/` directory with adapted specifications.
- [x] Define the Go-specific tool interface and command registry.
- [x] Implement a basic `Agent` wrapper that matches the expected Claude Code behavior.

### Phase 2: Enhanced TUI (`tview` implementation)
- [x] Expand `pkg/console/console.go` to support:
    - [x] Multiple view regions (Chat, Status, Context Usage).
    - [x] Rich text rendering (Markdown support in `tview`).
    - [x] Inline tool execution feedback.
    - [x] Keyboard shortcuts matching Claude Code (e.g., Ctrl+R for history).

### Phase 3: Essential Tools Porting
- [x] Port core tools to Go:
    - [x] `Read`, `Write`, `Edit` (File operations).
    - [x] `Bash` (Shell execution with background support).
    - [x] `Glob`, `Grep` (Search).
    - [x] `WebFetch`, `WebSearch` (Web access).

### Phase 4: Commands & Advanced Features
- [x] Implement slash commands:
    - [x] `/compact`, `/clear`, `/history`, `/help`.
    - [x] `/model`, `/config`.
- [x] Implement context management (compaction logic in Go).
- [x] Multi-agent delegation (using the `AgentTool` concept within `agentic`).

### Phase 5: Verification & Refinement
- [x] Add unit tests for tools and command parsing.
- [ ] Verify TUI performance and stability.
- [x] Finalize documentation and user guide.

## 4. Key Files to Create/Modify
- `spec_go/*.md`: Adapted specifications.
- `pkg/console/`: Main TUI logic and component wrappers.
- `pkg/tools/`: Tool implementations.
- `pkg/commands/`: Slash command handlers.
- `pkg/agent/`: Custom agent logic and prompt templates.
