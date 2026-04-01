# Claude Code (Go Port) — TUI Components & UX

## 1. TUI Architecture (`tview` implementation)
The TUI is organized as a set of nested flex containers managing various view regions.

### Main Layout:
- **Header**: Session title, current model, and global status.
- **ChatView** (`tview.TextView`):
    - Supports dynamic colors (using `tview` color tags).
    - Implements scrollback buffer management.
    - Renders Markdown-style formatting (bold, links, code blocks).
- **InputField** (`tview.InputField`):
    - Captures user prompts.
    - Supports basic command-line editing.
    - Integrates with a fuzzy-finding history picker (Ctrl+R).
- **StatusLine** (`tview.TextView`):
    - Shows token usage bar.
    - Displays current permission mode (e.g., Default, Auto).
    - Notifications area for background task updates.

## 2. Interactive Dialogs
For actions requiring user confirmation (e.g., tool execution in 'Ask Always' mode), the console will overlay modal dialogs.

### Types of Dialogs:
- **Permission Dialog**: Shows the tool name and proposed input. Options: `[Y]es`, `[N]o`, `[A]lways allow`, `[D]eny forever`.
- **Fuzzy Picker**: For history search (`Ctrl+R`) and file selection.
- **Config Editor**: Inline YAML/JSON editor for settings.

## 3. Keyboard Shortcuts
To maintain UX parity with Claude Code, the following shortcuts will be implemented:

| Shortcut | Action |
|---|---|
| `Enter` | Submit prompt / select option |
| `Ctrl+C` | Interrupt current agent run / cancel input |
| `Ctrl+R` | Open fuzzy history search |
| `Ctrl+O` | Toggle expanded message view |
| `Tab` | Autocomplete command or file path (if enabled) |
| `Esc` | Close modal dialog / clear input |

## 4. Rich Content Rendering
### Markdown Support:
- **Code Blocks**: Use distinct background colors and syntax highlighting if possible (via `chroma` or similar Go library integration).
- **Tool Invocations**: Rendered as distinct cards within the `ChatView` using border characters.
- **Images/PDFs**: Represented by placeholders/metatags with links to open in external viewers.

## 5. Background Task Monitoring
Tasks spawned in the background (e.g., long-running shell commands) will be listed in a sidebar or toggleable panel, showing their status and providing a way to "attach" to their output.
