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

package shell

import (
	engine "github.com/ZupIT/horusec-engine"

	"github.com/aledj02/horusx/internal/services/engines"
)

// NewRules returns a rule manager for Shell scripts. It intentionally ships no
// built-in rules: every Shell rule comes from the user-provided custom-rules
// JSON (--custom-rules-path). The manager exposes the shell extensions the
// engine should scan so the text engine applies the custom rules to shell files.
func NewRules() *engines.RuleManager {
	return engines.NewRuleManager(rules(), extensions())
}

func extensions() []string {
	return []string{".sh", ".bash", ".zsh", ".ksh", ".ksh93"}
}

func rules() []engine.Rule {
	return []engine.Rule{}
}
