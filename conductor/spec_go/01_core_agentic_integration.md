# Claude Code (Go Port) — Core Agentic Integration

## 1. Entry Point (`main.go`)
The `main.go` file serves as the bootstrap for the application.

### Startup Sequence:
1. **Config Loading**: Load YAML configuration from default or specified path.
2. **Registry Initialization**: Initialize the `agentic` registry with tools and agents.
3. **Auth Setup**: Configure JWT verifiers or API keys as needed.
4. **Console Launch**: Instantiate the `console.Launcher` and call `Run`.

## 2. Agentic Runner Implementation
The core interaction loop is managed by `runner.New` and its `Run` method.

### Agent Loop Logic:
- **Streaming Response**: Use `agent.StreamingModeSSE` for real-time output.
- **Event Handling**: Iterates over events from the runner:
    - **Content Events**: Append text to the TUI chat view.
    - **Tool Use Events**: Trigger tool execution logic.
    - **Final Response**: Record the full turn to history.

## 3. Tool Framework Alignment
Go tools must satisfy the `agent.Tool` interface.

```go
type Tool interface {
    Name() string
    Description() string
    Parameters() *genai.Schema
    Call(ctx context.Context, args map[string]any) (any, error)
}
```

### Adapted Tool Characteristics:
- **Permission Checking**: Integrated into the `Call` method or handled by a middleware layer.
- **Input Validation**: Leverages `genai.Schema` for structured input validation.
- **Progress Reporting**: Tools can emit status updates via a side-channel or specific return types if supported by the runner.

## 4. Context Management
Adapting the TypeScript `QueryEngine.ts` and `compact.ts` logic to Go.

### Compaction Strategy:
- **Token Tracking**: Track cumulative token usage using `Usage` data from the runner.
- **Automatic Compaction**: When usage hits ~90% of the model's limit, trigger a "compaction turn":
    - Send a system prompt instructing the model to summarize the history.
    - Replace the `session.Session` history with the summarized version.
    - Mark the compaction boundary in the transcript.

## 5. History & Persistence
- **Storage**: JSONL files in `~/.agentic/sessions/`.
- **Schema**: Align with the `SerializedMessage` and `Entry` types from the original spec for cross-compatibility if possible.
