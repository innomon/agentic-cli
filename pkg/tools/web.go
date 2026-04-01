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
	"io"
	"net/http"

	"google.golang.org/adk/tool"
	"google.golang.org/adk/tool/functiontool"
)

// WebFetchArgs defines the arguments for the WebFetchTool.
type WebFetchArgs struct {
	URL string `json:"url" description:"The URL to fetch."`
}

// WebFetchResult defines the result of the WebFetchTool.
type WebFetchResult struct {
	Content string `json:"content"`
}

// WebFetchHandler implements the web.Fetch logic.
func WebFetchHandler(ctx tool.Context, args WebFetchArgs) (WebFetchResult, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", args.URL, nil)
	if err != nil {
		return WebFetchResult{}, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return WebFetchResult{}, fmt.Errorf("failed to fetch URL: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return WebFetchResult{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return WebFetchResult{}, fmt.Errorf("failed to read response body: %w", err)
	}

	return WebFetchResult{Content: string(body)}, nil
}

// NewWebFetchTool creates a new WebFetch tool.
func NewWebFetchTool() (tool.Tool, error) {
	return functiontool.New(functiontool.Config{
		Name:        "web.Fetch",
		Description: "Retrieve the content of a URL.",
	}, WebFetchHandler)
}
