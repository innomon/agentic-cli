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
	"os"
	"path/filepath"
	"strings"

	"google.golang.org/adk/tool"
	"google.golang.org/adk/tool/functiontool"
)

// ReadArgs defines the arguments for the ReadTool.
type ReadArgs struct {
	Path string `json:"path" description:"The path to the file to read."`
}

// ReadResult defines the result of the ReadTool.
type ReadResult struct {
	Content string `json:"content"`
}

// ReadHandler implements the filesystem.Read logic.
func ReadHandler(ctx tool.Context, args ReadArgs) (ReadResult, error) {
	absPath, err := filepath.Abs(args.Path)
	if err != nil {
		return ReadResult{}, fmt.Errorf("invalid path: %w", err)
	}

	data, err := os.ReadFile(absPath)
	if err != nil {
		return ReadResult{}, fmt.Errorf("cannot read file %q: %w", args.Path, err)
	}

	return ReadResult{Content: string(data)}, nil
}

// NewReadTool creates a new Read tool.
func NewReadTool() (tool.Tool, error) {
	return functiontool.New(functiontool.Config{
		Name:        "filesystem.Read",
		Description: "Read the contents of a file.",
	}, ReadHandler)
}

// WriteArgs defines the arguments for the WriteTool.
type WriteArgs struct {
	Path    string `json:"path" description:"The path to the file to write."`
	Content string `json:"content" description:"The content to write to the file."`
}

// WriteResult defines the result of the WriteTool.
type WriteResult struct {
	Message string `json:"message"`
}

// WriteHandler implements the filesystem.Write logic.
func WriteHandler(ctx tool.Context, args WriteArgs) (WriteResult, error) {
	absPath, err := filepath.Abs(args.Path)
	if err != nil {
		return WriteResult{}, fmt.Errorf("invalid path: %w", err)
	}

	if err := os.MkdirAll(filepath.Dir(absPath), 0755); err != nil {
		return WriteResult{}, fmt.Errorf("failed to create directories: %w", err)
	}

	if err := os.WriteFile(absPath, []byte(args.Content), 0644); err != nil {
		return WriteResult{}, fmt.Errorf("cannot write file %q: %w", args.Path, err)
	}

	return WriteResult{Message: fmt.Sprintf("Successfully wrote to %s", args.Path)}, nil
}

// NewWriteTool creates a new Write tool.
func NewWriteTool() (tool.Tool, error) {
	return functiontool.New(functiontool.Config{
		Name:        "filesystem.Write",
		Description: "Create or overwrite a file with the specified content.",
	}, WriteHandler)
}

// EditArgs defines the arguments for the EditTool.
type EditArgs struct {
	Path      string `json:"path" description:"The path to the file to edit."`
	OldString string `json:"old_string" description:"The exact literal text to replace."`
	NewString string `json:"new_string" description:"The text to replace 'old_string' with."`
}

// EditResult defines the result of the EditTool.
type EditResult struct {
	Message string `json:"message"`
}

// EditHandler implements the filesystem.Edit logic.
func EditHandler(ctx tool.Context, args EditArgs) (EditResult, error) {
	absPath, err := filepath.Abs(args.Path)
	if err != nil {
		return EditResult{}, fmt.Errorf("invalid path: %w", err)
	}

	data, err := os.ReadFile(absPath)
	if err != nil {
		return EditResult{}, fmt.Errorf("cannot read file %q: %w", args.Path, err)
	}

	content := string(data)
	if !strings.Contains(content, args.OldString) {
		return EditResult{}, fmt.Errorf("old_string not found in file %q", args.Path)
	}

	if strings.Count(content, args.OldString) > 1 {
		return EditResult{}, fmt.Errorf("old_string is ambiguous (found multiple occurrences) in file %q", args.Path)
	}

	newContent := strings.Replace(content, args.OldString, args.NewString, 1)

	if err := os.WriteFile(absPath, []byte(newContent), 0644); err != nil {
		return EditResult{}, fmt.Errorf("failed to write edited content to %q: %w", args.Path, err)
	}

	return EditResult{Message: fmt.Sprintf("Successfully edited %s", args.Path)}, nil
}

// NewEditTool creates a new Edit tool.
func NewEditTool() (tool.Tool, error) {
	return functiontool.New(functiontool.Config{
		Name:        "filesystem.Edit",
		Description: "Replace a specific block of text in a file with new content.",
	}, EditHandler)
}
