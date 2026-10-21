// Copyright 2021 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
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

package images

import "github.com/ZupIT/horusec-devkit/pkg/enums/languages"

const (
	DefaultRegistry = "docker.io"
	C               = "aledj02/horusx-c:v2.12.0"
	Csharp          = "aledj02/horusx-csharp:v2.12.0"
	Elixir          = "aledj02/horusx-elixir:v2.12.0"
	Generic         = "aledj02/horusx-generic:v2.12.2"
	Go              = "aledj02/horusx-go:v2.10.1"
	HCL             = "aledj02/horusx-hcl:v2.12.0"
	Javascript      = "aledj02/horusx-js:v2.10.1"
	Leaks           = "aledj02/horusx-leaks:v2.12.0"
	PHP             = "aledj02/horusx-php:v2.12.0"
	Python          = "aledj02/horusx-python:v2.10.1"
	Ruby            = "aledj02/horusx-ruby:v2.12.0"
	Shell           = "aledj02/horusx-shell:v2.10.1"
)

func MapValues() map[languages.Language]string {
	return map[languages.Language]string{
		languages.CSharp:     Csharp,
		languages.Leaks:      Leaks,
		languages.Go:         Go,
		languages.Javascript: Javascript,
		languages.Python:     Python,
		languages.Ruby:       Ruby,
		languages.HCL:        HCL,
		languages.Generic:    Generic,
		languages.PHP:        PHP,
		languages.Elixir:     Elixir,
		languages.Shell:      Shell,
		languages.C:          C,
	}
}
