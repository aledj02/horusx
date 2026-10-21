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

package horusecgo

import (
	"github.com/ZupIT/horusec-devkit/pkg/enums/languages"

	goengine "github.com/ZupIT/horusec/internal/services/engines/go"
	"github.com/ZupIT/horusec/internal/services/formatters"
)

// NewFormatter returns a DefaultFormatter for Go that runs the user's
// custom-rules JSON (loaded from --custom-rules-path) against .go files. It is
// wired into analyzer.runner alongside the existing Docker-backed gosec and
// nancy tools.
func NewFormatter(service formatters.IService) formatters.IFormatter {
	return formatters.NewDefaultFormatter(service, goengine.NewRules(), languages.Go)
}
