# Claude Code (Go Port) — Spec Index

This index covers the adapted specifications for implementing Claude Code features within the Go-based `agentic` framework.

---

## Spec Files

| # | File | What's Inside |
|---|------|---------------|
| 00 | [00_overview.md](00_overview.md) | Master architecture, Go frameworks, and data flow. |
| 01 | [01_core_agentic_integration.md](01_core_agentic_integration.md) | `main.go` bootstrap, agent runner loop, and context management. |
| 02 | [02_tui_components.md](02_tui_components.md) | `tview` layout, interactive dialogs, and keyboard shortcuts. |
| 03 | [03_tools_go_implementation.md](03_tools_go_implementation.md) | Porting core tools (FS, Shell, Web, Search) to Go. |

## Quick Comparison

| Feature | Original (TS/Rust) | Go Port |
|---|---|---|
| Core Engine | Node.js / Bun / Rust | Go (Native) |
| Agent Logic | Custom `QueryEngine.ts` | ADK `runner.Runner` |
| UI | React / Ink | `tview` / `tcell` |
| Tools | Class-based `Tool.ts` | Interface-based `agent.Tool` |
| Persistence | `sessionStorage.ts` | JSONL in `~/.agentic` |
