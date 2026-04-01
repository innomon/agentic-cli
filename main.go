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

package main

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/innomon/agentic/pkg/auth"
	"github.com/innomon/agentic/pkg/config"
	"agenticli/pkg/console"
	"github.com/innomon/agentic/pkg/registry"
	_ "github.com/innomon/agentic/pkg/gnogent"
	_ "github.com/innomon/agentic/pkg/gomlx"
	_ "github.com/innomon/agentic/pkg/prologmem"
	_ "github.com/innomon/agentic/pkg/routing"
	_ "github.com/innomon/agentic/pkg/wasm"

	"github.com/a2aproject/a2a-go/a2asrv"
)

func main() {
	ctx := context.Background()
	var cfg *config.Config
	var err error
	var largs = 1

	// Check if the first argument is a config file
	if len(os.Args) > 1 && (strings.HasSuffix(os.Args[1], ".yml") || strings.HasSuffix(os.Args[1], ".yaml")) {
		cfg, err = config.Load(os.Args[1])
		if err != nil {
			log.Fatalf("Failed to load config: %v", err)
		}
		largs = 2
	} else {
		cfg, err = config.LoadDefault()
		if err != nil {
			log.Fatalf("Failed to load config: %v", err)
		}
	}

	reg := registry.New(cfg)
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered from panic: %v", r)
		}
		if err := reg.Close(); err != nil {
			log.Printf("Error closing registry: %v", err)
		}
	}()

	launcherConfig, err := reg.BuildLauncherConfig(ctx)
	if err != nil {
		log.Fatalf("Failed to build launcher config: %v", err)
	}

	if authCfg := reg.Config().Auth; authCfg != nil && authCfg.JWT != nil {
		jwt := authCfg.JWT
		verifier, err := auth.NewJWTVerifier(jwt.PublicKeyPath, jwt.Issuer, jwt.Audience)
		if err != nil {
			log.Fatalf("Failed to create JWT verifier: %v", err)
		}
		launcherConfig.A2AOptions = append(launcherConfig.A2AOptions, a2asrv.WithCallInterceptor(&auth.JWTInterceptor{Verifier: verifier}))
		log.Printf("JWT authentication enabled (issuer=%s, audience=%s)", jwt.Issuer, jwt.Audience)
	}

	// Always use the console launcher
	l := console.New()
	
	// Parse remaining flags if any
	if _, err := l.Parse(os.Args[largs:]); err != nil {
		log.Fatalf("Failed to parse console flags: %v", err)
	}

	if err := l.Run(ctx, launcherConfig); err != nil {
		log.Fatalf("Console error: %v", err)
	}
}
