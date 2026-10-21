// Copyright 2026 HORUSX
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

// Package customrules_test holds end-to-end tests that load the real
// custom-rules JSON files from custom_rules/ and run the horusec-engine text
// scanner against the vulnerable fixtures in examples/.
//
// These tests are NOT Docker-dependent and exercise the full path:
//
//	custom rule JSON → CustomRule → text.Rule → engine.Engine.Run → Finding
package customrules_test

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/ZupIT/horusec-devkit/pkg/enums/languages"
	engine "github.com/ZupIT/horusec-engine"
	"github.com/ZupIT/horusec-engine/text"
	"github.com/stretchr/testify/require"

	"github.com/ZupIT/horusec/config"
	customrules "github.com/ZupIT/horusec/internal/services/custom_rules"
)

// rootPath goes three levels up from this test package to the horusec repo root.
const rootPath = "../../../"

type e2eCase struct {
	name         string
	rulesJSON    string
	ruleID       string
	fixturePath  string
	lang         languages.Language
	findingCount int
}

func TestE2E_CustormRulesDischargeFixture(t *testing.T) {
	cases := []e2eCase{
		{
			name:         "Python-HS-PYTHON-100-pickle_loads",
			rulesJSON:    filepath.Join(rootPath, "custom_rules", "python.json"),
			ruleID:       "HS-PYTHON-100",
			fixturePath:  filepath.Join(rootPath, "examples", "python", "example1", "pickle_loads.py"),
			lang:         languages.Python,
			findingCount: 1,
		},
		{
			name:         "Go-HS-GO-100-insecure_TLS",
			rulesJSON:    filepath.Join(rootPath, "custom_rules", "go.json"),
			ruleID:       "HS-GO-100",
			fixturePath:  filepath.Join(rootPath, "examples", "go", "example1", "insecure_tls.go"),
			lang:         languages.Go,
			findingCount: 1,
		},
		{
			name:         "Shell-HS-SHELL-100-install_from_url",
			rulesJSON:    filepath.Join(rootPath, "custom_rules", "shell.json"),
			ruleID:       "HS-SHELL-100",
			fixturePath:  filepath.Join(rootPath, "examples", "shell", "example1", "install_from_url.sh"),
			lang:         languages.Shell,
			findingCount: 1,
		},
		{
			name:         "Swift-HS-SWIFT-101-unarchive",
			rulesJSON:    filepath.Join(rootPath, "custom_rules", "swift.json"),
			ruleID:       "HS-SWIFT-101",
			fixturePath:  filepath.Join(rootPath, "examples", "swift", "example1", "insecure_unarchive.swift"),
			lang:         languages.Swift,
			findingCount: 1,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			cfg := config.New()
			cfg.CustomRulesPath = tt.rulesJSON

			svc := customrules.NewCustomRulesService(cfg)
			rules := svc.Load(tt.lang)

			require.NotEmpty(t, rules,
				"expected at least 1 rule loaded from %s for language %s", tt.rulesJSON, tt.lang)

			var rule *text.Rule
			for _, r := range rules {
				if tr, ok := r.(*text.Rule); ok && tr.ID == tt.ruleID {
					rule = tr
					break
				}
			}
			require.NotNil(t, rule, "rule %s not found in %s", tt.ruleID, tt.rulesJSON)

			eng := engine.NewEngine(0, "*")
			findings, err := eng.Run(context.Background(), tt.fixturePath, rule)
			require.NoError(t, err)
			require.GreaterOrEqual(t, len(findings), tt.findingCount,
				"expected >= %d findings for fixture %s, got %d", tt.findingCount, tt.fixturePath, len(findings))

			for i, f := range findings {
				t.Logf("finding[%d] id=%s line=%d sample=%q", i, f.ID, f.SourceLocation.Line, f.CodeSample)
			}
		})
	}
}
