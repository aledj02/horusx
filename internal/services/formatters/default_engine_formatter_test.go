// Copyright 2020 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
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

package formatters_test

import (
	"testing"

	"github.com/ZupIT/horusec-devkit/pkg/entities/analysis"
	engine "github.com/ZupIT/horusec-engine"
	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec/internal/services/formatters"
	"github.com/ZupIT/horusec/internal/services/formatters/csharp/horusxcsharp"
	"github.com/ZupIT/horusec/internal/services/formatters/dart/horusxdart"
	"github.com/ZupIT/horusec/internal/services/formatters/go/horusxgo"
	"github.com/ZupIT/horusec/internal/services/formatters/java/horusxjava"
	"github.com/ZupIT/horusec/internal/services/formatters/javascript/horusxjavascript"
	"github.com/ZupIT/horusec/internal/services/formatters/kotlin/horusxkotlin"
	"github.com/ZupIT/horusec/internal/services/formatters/leaks/horusxleaks"
	"github.com/ZupIT/horusec/internal/services/formatters/nginx/horusxnginx"
	"github.com/ZupIT/horusec/internal/services/formatters/python/horusxpython"
	"github.com/ZupIT/horusec/internal/services/formatters/shell/horusxshell"
	"github.com/ZupIT/horusec/internal/services/formatters/swift/horusxswift"
	"github.com/ZupIT/horusec/internal/services/formatters/yaml/horusxkubernetes"
	"github.com/ZupIT/horusec/internal/utils/testutil"
)

func TestStartAnalysis(t *testing.T) {
	testcases := []struct {
		engine    string
		formatter func(formatters.IService) formatters.IFormatter
	}{
		{
			engine:    "Csharp",
			formatter: horusxcsharp.NewFormatter,
		},
		{
			engine:    "Dart",
			formatter: horusxdart.NewFormatter,
		},
		{
			engine:    "Java",
			formatter: horusxjava.NewFormatter,
		},
		{
			engine:    "Javascript",
			formatter: horusxjavascript.NewFormatter,
		},
		{
			engine:    "Kotlin",
			formatter: horusxkotlin.NewFormatter,
		},
		{
			engine:    "Leaks",
			formatter: horusxleaks.NewFormatter,
		},
		{
			engine:    "Nginx",
			formatter: horusxnginx.NewFormatter,
		},
		{
			engine:    "Kubernetes",
			formatter: horusxkubernetes.NewFormatter,
		},
		{
			engine:    "Python",
			formatter: horusxpython.NewFormatter,
		},
		{
			engine:    "Go",
			formatter: horusxgo.NewFormatter,
		},
		{
			engine:    "Shell",
			formatter: horusxshell.NewFormatter,
		},
		{
			engine:    "Swift",
			formatter: horusxswift.NewFormatter,
		},
	}

	for _, tt := range testcases {
		t.Run(tt.engine, func(t *testing.T) {
			t.Run("should success execute analysis without errors", func(t *testing.T) {
				analysis := &analysis.Analysis{}
				service := testutil.NewFormatterMock()

				service.On("LogDebugWithReplace")
				service.On("SetToolFinishedAnalysis")
				service.On("SetAnalysisError")
				service.On("ToolIsToIgnore").Return(false)
				service.On("GetConfigProjectPath").Return(".")
				service.On("ParseFindingsToVulnerabilities").Return(nil)
				service.On("GetCustomRulesByLanguage").Return([]engine.Rule{})

				assert.NotPanics(t, func() {
					tt.formatter(service).StartAnalysis("")
				})

				assert.Empty(t, len(analysis.Errors))
			})

			t.Run("should return error when getting text unit", func(t *testing.T) {
				service := testutil.NewFormatterMock()

				service.On("LogDebugWithReplace")
				service.On("SetToolFinishedAnalysis")
				service.On("SetAnalysisError")
				service.On("ToolIsToIgnore").Return(false)
				service.On("GetConfigProjectPath").Return(".")
				service.On("ParseFindingsToVulnerabilities").Return(nil)
				service.On("GetCustomRulesByLanguage").Return([]engine.Rule{})

				assert.NotPanics(t, func() {
					tt.formatter(service).StartAnalysis("")
				})
			})

			t.Run("should ignore this tool", func(t *testing.T) {
				service := testutil.NewFormatterMock()

				service.On("ToolIsToIgnore").Return(true)

				assert.NotPanics(t, func() {
					tt.formatter(service).StartAnalysis("")
				})
			})
		})
	}
}
