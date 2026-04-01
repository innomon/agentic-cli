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

package commands

import (
	"context"
	"fmt"
	"strings"
)

// RegisterDefaultCommands adds the standard set of slash commands to the registry.
func RegisterDefaultCommands(r *Registry) {
	r.Register(Command{
		Name:        "help",
		Description: "Show available commands and their descriptions",
		Handler: func(ctx context.Context, args []string) (string, error) {
			var b strings.Builder
			b.WriteString("Available Commands:\n")
			for _, cmd := range r.GetCommands() {
				b.WriteString(fmt.Sprintf("/%-10s - %s\n", cmd.Name, cmd.Description))
			}
			return b.String(), nil
		},
	})

	r.Register(Command{
		Name:        "clear",
		Description: "Clear the chat history view",
		Handler: func(ctx context.Context, args []string) (string, error) {
			// This will be handled by the TUI by returning a special signal if needed,
			// or we can just return a message and let the caller clear the view.
			return "CLEAR_SIGNAL", nil
		},
	})

	r.Register(Command{
		Name:        "history",
		Description: "Show recent command history",
		Handler: func(ctx context.Context, args []string) (string, error) {
			// In a real app, this would access the history from somewhere.
			// For now, we'll just return a placeholder.
			return "History functionality is available via Ctrl+R.", nil
		},
	})

	r.Register(Command{
		Name:        "exit",
		Description: "Quit the application",
		Handler: func(ctx context.Context, args []string) (string, error) {
			return "Quitting...", nil
		},
	})

	r.Register(Command{
		Name:        "quit",
		Description: "Quit the application",
		Handler: func(ctx context.Context, args []string) (string, error) {
			return "Quitting...", nil
		},
	})
}
