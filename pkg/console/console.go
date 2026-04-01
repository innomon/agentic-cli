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

package console

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"agenticli/pkg/agent"
	"agenticli/pkg/commands"

	adkagent "google.golang.org/adk/agent"
	"google.golang.org/adk/cmd/launcher"
	"google.golang.org/genai"
)

// Launcher implements a TUI console.
type Launcher struct {
	flags    *flag.FlagSet
	config   *config
	commands *commands.Registry
	history  []string
	hIndex   int
}

type config struct {
	streamingMode       adkagent.StreamingMode
	streamingModeString string
}

// New creates a new TUI console launcher.
func New() *Launcher {
	cfg := &config{}
	fs := flag.NewFlagSet("console", flag.ContinueOnError)
	fs.StringVar(&cfg.streamingModeString, "streaming_mode", string(adkagent.StreamingModeSSE),
		"streaming mode (none|sse)")

	registry := commands.NewRegistry()
	commands.RegisterDefaultCommands(registry)

	return &Launcher{
		config:   cfg,
		flags:    fs,
		commands: registry,
		history:  []string{},
		hIndex:   -1,
	}
}

// Keyword returns the command keyword for this launcher.
func (l *Launcher) Keyword() string {
	return "console"
}

// CommandLineSyntax returns usage documentation.
func (l *Launcher) CommandLineSyntax() string {
	return `TUI Console mode with file attachment support.

Usage: ./agentic console [options]

Attach files using @/path/to/file syntax.
/exit to quit.
`
}

// SimpleDescription returns a short description.
func (l *Launcher) SimpleDescription() string {
	return "TUI interactive console with file attachment support"
}

// Parse parses command-line arguments.
func (l *Launcher) Parse(args []string) ([]string, error) {
	if err := l.flags.Parse(args); err != nil {
		return nil, fmt.Errorf("failed to parse flags: %w", err)
	}
	mode := l.config.streamingModeString
	if mode != string(adkagent.StreamingModeNone) && mode != string(adkagent.StreamingModeSSE) {
		return nil, fmt.Errorf("invalid streaming_mode: %s", mode)
	}
	l.config.streamingMode = adkagent.StreamingMode(mode)
	return l.flags.Args(), nil
}

// Run executes the TUI console.
func (l *Launcher) Run(ctx context.Context, cfg *launcher.Config) error {
	const (
		userID = "console_user"
	)

	agentWrapper, err := agent.New(ctx, cfg, userID)
	if err != nil {
		return err
	}

	// TUI setup
	app := tview.NewApplication()

	// Chat history view
	chatView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWordWrap(true).
		SetChangedFunc(func() {
			app.Draw()
		})
	chatView.SetBorder(true).SetTitle(" Chat History ")

	// Input field
	inputField := tview.NewInputField().
		SetLabel("User -> ").
		SetFieldWidth(0)
	inputField.SetBorder(true).SetTitle(" Input ")

	// Status and metadata view
	statusView := tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignLeft)
	statusView.SetText(" [yellow]Commands:[white] /help, /exit. [yellow]History:[white] Up/Down. [yellow]Attach:[white] @file")

	// Context and token usage view (placeholder)
	usageView := tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignRight)
	usageView.SetText(" [green]Session:[white] " + agentWrapper.GetSessionID()[:8] + "... ")

	// Top bar for metadata and usage
	topBar := tview.NewFlex().
		AddItem(statusView, 0, 1, false).
		AddItem(usageView, 30, 0, false)

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(chatView, 0, 1, false).
		AddItem(inputField, 3, 1, true).
		AddItem(topBar, 1, 1, false)

	filePattern := regexp.MustCompile(`@([^\s]+)`)

	fmt.Fprintf(chatView, "[green]Welcome to AgentiCli![white]\n")
	fmt.Fprintf(chatView, "Attach files using @/path/to/file syntax.\n")
	fmt.Fprintf(chatView, "Type /help for commands, /exit to quit.\n\n")

	// History navigation
	inputField.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyUp {
			if l.hIndex < len(l.history)-1 {
				l.hIndex++
				inputField.SetText(l.history[len(l.history)-1-l.hIndex])
			}
			return nil
		}
		if event.Key() == tcell.KeyDown {
			if l.hIndex > 0 {
				l.hIndex--
				inputField.SetText(l.history[len(l.history)-1-l.hIndex])
			} else if l.hIndex == 0 {
				l.hIndex = -1
				inputField.SetText("")
			}
			return nil
		}
		if event.Key() == tcell.KeyCtrlR {
			if len(l.history) == 0 {
				return nil
			}
			list := tview.NewList()
			for i := len(l.history) - 1; i >= 0; i-- {
				list.AddItem(l.history[i], "", 0, nil)
			}
			list.SetSelectedFunc(func(i int, mainText, secondaryText string, shortcut rune) {
				inputField.SetText(mainText)
				app.SetRoot(flex, true)
				app.SetFocus(inputField)
			})
			list.SetDoneFunc(func() {
				app.SetRoot(flex, true)
				app.SetFocus(inputField)
			})
			list.SetBorder(true).SetTitle(" Search History ")
			app.SetRoot(list, true)
			return nil
		}
		return event
	})

	// Input handling
	inputField.SetDoneFunc(func(key tcell.Key) {
		if key != tcell.KeyEnter {
			return
		}

		input := strings.TrimSpace(inputField.GetText())
		if input == "" {
			return
		}

		// Add to history
		if len(l.history) == 0 || l.history[len(l.history)-1] != input {
			l.history = append(l.history, input)
		}
		l.hIndex = -1

		inputField.SetText("")

		// Handle slash commands
		output, handled, err := l.commands.Execute(ctx, input)
		if handled {
			if err != nil {
				fmt.Fprintf(chatView, "[red]Command Error: %v[white]\n", err)
			} else if output == "CLEAR_SIGNAL" {
				chatView.Clear()
			} else if output != "" {
				fmt.Fprintf(chatView, "[yellow]%s[white]\n", output)
			}
			if input == "/exit" || input == "/quit" {
				app.Stop()
			}
			return
		}

		fmt.Fprintf(chatView, "[blue]User -> [white]%s\n", input)

		userMsg, err := parseInputTUI(input, filePattern, chatView)
		if err != nil {
			fmt.Fprintf(chatView, "[red]Error: %v[white]\n", err)
			return
		}

		go func() {
			fmt.Fprintf(chatView, "[green]Agent -> [white]")
			
			onToolUse := func(toolName string, args map[string]any) {
				app.QueueUpdateDraw(func() {
					fmt.Fprintf(chatView, "\n[yellow]Executing tool: [white]%s(args: %v)\n", toolName, args)
				})
			}

			// Use the wrapper to run the agent
			err := agentWrapper.Run(ctx, userMsg, chatView, onToolUse)
			if err != nil {
				app.QueueUpdateDraw(func() {
					fmt.Fprintf(chatView, "\n[red]Error: %v[white]\n", err)
				})
			}
			app.QueueUpdateDraw(func() {
				fmt.Fprintf(chatView, "\n\n")
			})
		}()
	})

	if err := app.SetRoot(flex, true).Run(); err != nil {
		return err
	}

	return nil
}

func parseInputTUI(input string, filePattern *regexp.Regexp, chatView *tview.TextView) (*genai.Content, error) {
	matches := filePattern.FindAllStringSubmatch(input, -1)
	textContent := strings.TrimSpace(filePattern.ReplaceAllString(input, ""))

	var parts []*genai.Part

	if textContent != "" {
		parts = append(parts, genai.NewPartFromText(textContent))
	}

	for _, match := range matches {
		filePath := match[1]

		if strings.HasPrefix(filePath, "~") {
			home, err := os.UserHomeDir()
			if err == nil {
				filePath = filepath.Join(home, filePath[1:])
			}
		}

		absPath, err := filepath.Abs(filePath)
		if err != nil {
			return nil, fmt.Errorf("invalid path %q: %w", filePath, err)
		}

		data, err := os.ReadFile(absPath)
		if err != nil {
			return nil, fmt.Errorf("cannot read %q: %w", filePath, err)
		}

		mimeType := getMIMEType(filePath)
		parts = append(parts, &genai.Part{
			InlineData: &genai.Blob{
				MIMEType: mimeType,
				Data:     data,
			},
		})

		fmt.Fprintf(chatView, "[yellow][Attached: %s (%s, %d bytes)][white]\n", filepath.Base(filePath), mimeType, len(data))
	}

	if len(parts) == 0 {
		return nil, fmt.Errorf("empty message")
	}

	return &genai.Content{
		Role:  genai.RoleUser,
		Parts: parts,
	}, nil
}

func getMIMEType(path string) string {
	ext := strings.ToLower(filepath.Ext(path))
	types := map[string]string{
		".pdf":  "application/pdf",
		".png":  "image/png",
		".jpg":  "image/jpeg",
		".jpeg": "image/jpeg",
		".gif":  "image/gif",
		".webp": "image/webp",
		".txt":  "text/plain",
		".json": "application/json",
		".csv":  "text/csv",
	}
	if t, ok := types[ext]; ok {
		return t
	}
	return "application/octet-stream"
}
