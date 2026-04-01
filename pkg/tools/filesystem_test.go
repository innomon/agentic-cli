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
	"os"
	"path/filepath"
	"testing"

	"google.golang.org/adk/tool"
)

// mockContext implements tool.Context for testing.
type mockContext struct {
	tool.Context
}

func (m *mockContext) Done() <-chan struct{} {
	return nil
}

func TestFilesystemHandlers(t *testing.T) {
	tempDir := t.TempDir()
	filePath := filepath.Join(tempDir, "test.txt")
	content := "Hello, World!"

	ctx := &mockContext{}
	
	// Test WriteHandler
	writeResp, err := WriteHandler(ctx, WriteArgs{
		Path:    filePath,
		Content: content,
	})
	if err != nil {
		t.Fatalf("WriteHandler failed: %v", err)
	}
	if writeResp.Message != "Successfully wrote to "+filePath {
		t.Fatalf("Unexpected WriteHandler response: %v", writeResp.Message)
	}

	// Verify file was written
	data, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Failed to read back file: %v", err)
	}
	if string(data) != content {
		t.Fatalf("File content mismatch: expected %q, got %q", content, string(data))
	}

	// Test ReadHandler
	readResp, err := ReadHandler(ctx, ReadArgs{
		Path: filePath,
	})
	if err != nil {
		t.Fatalf("ReadHandler failed: %v", err)
	}
	if readResp.Content != content {
		t.Fatalf("ReadHandler response mismatch: expected %q, got %q", content, readResp.Content)
	}

	// Test EditHandler
	newContent := "Hello, Agentic!"
	editResp, err := EditHandler(ctx, EditArgs{
		Path:       filePath,
		OldString: "World",
		NewString: "Agentic",
	})
	if err != nil {
		t.Fatalf("EditHandler failed: %v", err)
	}
	if editResp.Message != "Successfully edited "+filePath {
		t.Fatalf("Unexpected EditHandler response: %v", editResp.Message)
	}
	
	data, err = os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Failed to read back edited file: %v", err)
	}
	if string(data) != newContent {
		t.Fatalf("Edited file content mismatch: expected %q, got %q", newContent, string(data))
	}
}
