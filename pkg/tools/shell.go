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

package tools

import (
	"fmt"
	"os/exec"
	"runtime"

	"google.golang.org/adk/tool"
	"google.golang.org/adk/tool/functiontool"
)

// BashArgs defines the arguments for the BashTool.
type BashArgs struct {
	Command string `json:"command" description:"The command to execute."`
}

// BashResult defines the result of the BashTool.
type BashResult struct {
	Output string `json:"output"`
}

// BashHandler implements the shell.Bash logic.
func BashHandler(ctx tool.Context, args BashArgs) (BashResult, error) {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.CommandContext(ctx, "cmd", "/c", args.Command)
	} else {
		cmd = exec.CommandContext(ctx, "bash", "-c", args.Command)
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		return BashResult{Output: string(output)}, fmt.Errorf("command execution failed: %w", err)
	}

	return BashResult{Output: string(output)}, nil
}

// NewBashTool creates a new Bash tool.
func NewBashTool() (tool.Tool, error) {
	return functiontool.New(functiontool.Config{
		Name:        "shell.Bash",
		Description: "Execute a bash command.",
	}, BashHandler)
}
