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

package horusxshell

import (
	"github.com/ZupIT/horusec-devkit/pkg/enums/languages"

	"github.com/aledj02/horusx/internal/services/engines/shell"
	"github.com/aledj02/horusx/internal/services/formatters"
)

// NewFormatter returns a DefaultFormatter for Shell scripts that runs the
// user's custom-rules JSON (loaded from --custom-rules-path) against shell
// files. It is wired into analyzer.runner alongside the existing Docker-backed
// shellcheck tool.
func NewFormatter(service formatters.IService) formatters.IFormatter {
	return formatters.NewDefaultFormatter(service, shell.NewRules(), languages.Shell)
}
