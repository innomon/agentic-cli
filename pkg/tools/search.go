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
	"regexp"
	"strings"

	"google.golang.org/adk/tool"
	"google.golang.org/adk/tool/functiontool"
)

// GlobArgs defines the arguments for the GlobTool.
type GlobArgs struct {
	Pattern string `json:"pattern" description:"The glob pattern to match against."`
}

// GlobResult defines the result of the GlobTool.
type GlobResult struct {
	Files []string `json:"files"`
}

// GlobHandler implements the search.Glob logic.
func GlobHandler(ctx tool.Context, args GlobArgs) (GlobResult, error) {
	var files []string
	if strings.Contains(args.Pattern, "**") {
		baseDir := "."
		err := filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				files = append(files, path)
			}
			return nil
		})
		if err != nil {
			return GlobResult{}, err
		}
	} else {
		var err error
		files, err = filepath.Glob(args.Pattern)
		if err != nil {
			return GlobResult{}, err
		}
	}

	return GlobResult{Files: files}, nil
}

// NewGlobTool creates a new Glob tool.
func NewGlobTool() (tool.Tool, error) {
	return functiontool.New(functiontool.Config{
		Name:        "search.Glob",
		Description: "List files matching a glob pattern.",
	}, GlobHandler)
}

// GrepArgs defines the arguments for the GrepTool.
type GrepArgs struct {
	Pattern string `json:"pattern" description:"The regex pattern to search for."`
	Include string `json:"include,omitempty" description:"Glob pattern for files to include in search."`
}

// GrepResult defines the result of the GrepTool.
type GrepResult struct {
	Results string `json:"results"`
}

// GrepHandler implements the search.Grep logic.
func GrepHandler(ctx tool.Context, args GrepArgs) (GrepResult, error) {
	re, err := regexp.Compile(args.Pattern)
	if err != nil {
		return GrepResult{}, fmt.Errorf("invalid regex: %w", err)
	}

	include := "*"
	if args.Include != "" {
		include = args.Include
	}

	files, err := filepath.Glob(include)
	if err != nil {
		return GrepResult{}, err
	}

	var results []string
	for _, file := range files {
		data, err := os.ReadFile(file)
		if err != nil {
			continue
		}

		lines := strings.Split(string(data), "\n")
		for i, line := range lines {
			if re.MatchString(line) {
				results = append(results, fmt.Sprintf("%s:%d: %s", file, i+1, line))
			}
		}
	}

	return GrepResult{Results: strings.Join(results, "\n")}, nil
}

// NewGrepTool creates a new Grep tool.
func NewGrepTool() (tool.Tool, error) {
	return functiontool.New(functiontool.Config{
		Name:        "search.Grep",
		Description: "Search for a pattern in files.",
	}, GrepHandler)
}
