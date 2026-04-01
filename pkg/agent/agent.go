// Copyright 2026 Innomon
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package agent

import (
	"context"
	"fmt"
	"io"

	"google.golang.org/adk/agent"
	"google.golang.org/adk/cmd/launcher"
	"google.golang.org/adk/runner"
	"google.golang.org/adk/session"
	"google.golang.org/genai"
)

// Wrapper provides a high-level interface to the underlying ADK runner.
type Wrapper struct {
	runner  *runner.Runner
	session session.Session
	userID  string
}

// New creates a new agent wrapper.
func New(ctx context.Context, cfg *launcher.Config, userID string) (*Wrapper, error) {
	const appName = "AgentiCli"

	sessionService := cfg.SessionService
	if sessionService == nil {
		sessionService = session.InMemoryService()
	}

	resp, err := sessionService.Create(ctx, &session.CreateRequest{
		AppName: appName,
		UserID:  userID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create session: %w", err)
	}

	rootAgent := cfg.AgentLoader.RootAgent()
	r, err := runner.New(runner.Config{
		AppName:         appName,
		Agent:           rootAgent,
		SessionService:  sessionService,
		ArtifactService: cfg.ArtifactService,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create runner: %w", err)
	}

	return &Wrapper{
		runner:  r,
		session: resp.Session,
		userID:  userID,
	}, nil
}

// Run executes the agent on the given input and streams events back.
func (w *Wrapper) Run(ctx context.Context, input *genai.Content, w_out io.Writer, onToolUse func(toolName string, args map[string]any)) error {
	streamingMode := agent.StreamingModeSSE
	
	for event, err := range w.runner.Run(ctx, w.userID, w.session.ID(), input, agent.RunConfig{
		StreamingMode: streamingMode,
	}) {
		if err != nil {
			return err
		}
		
		if event.Content != nil && onToolUse != nil {
			for _, p := range event.Content.Parts {
				if p.FunctionCall != nil {
					onToolUse(p.FunctionCall.Name, p.FunctionCall.Args)
				}
			}
		}

		if event.Content != nil {
			for _, p := range event.Content.Parts {
				if p.Text != "" {
					fmt.Fprint(w_out, p.Text)
				}
			}
		}
	}
	return nil
}

// GetSessionID returns the current session ID.
func (w *Wrapper) GetSessionID() string {
	return w.session.ID()
}
