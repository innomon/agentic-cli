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

// Command represents a slash command.
type Command struct {
	Name        string
	Description string
	Handler     func(ctx context.Context, args []string) (string, error)
}

// Registry manages the set of available slash commands.
type Registry struct {
	commands map[string]Command
}

// NewRegistry creates a new command registry.
func NewRegistry() *Registry {
	return &Registry{
		commands: make(map[string]Command),
	}
}

// Register adds a new command to the registry.
func (r *Registry) Register(cmd Command) {
	r.commands[cmd.Name] = cmd
}

// Execute looks up and runs a command by name.
func (r *Registry) Execute(ctx context.Context, input string) (string, bool, error) {
	if !strings.HasPrefix(input, "/") {
		return "", false, nil
	}

	parts := strings.Fields(input)
	if len(parts) == 0 {
		return "", false, nil
	}

	cmdName := parts[0][1:] // Remove the leading slash
	cmd, ok := r.commands[cmdName]
	if !ok {
		return "", true, fmt.Errorf("unknown command: /%s", cmdName)
	}

	output, err := cmd.Handler(ctx, parts[1:])
	return output, true, err
}

// GetCommands returns all registered commands.
func (r *Registry) GetCommands() []Command {
	var cmds []Command
	for _, cmd := range r.commands {
		cmds = append(cmds, cmd)
	}
	return cmds
}
