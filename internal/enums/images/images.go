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
	C               = "horuszup/horusx-c:v1.0.1"
	Csharp          = "horuszup/horusx-csharp:v1.2.0"
	Elixir          = "horuszup/horusx-elixir:v1.1.0"
	Generic         = "horuszup/horusx-generic:v1.2.0"
	Go              = "aledj02/horusx-go:v2.10.1"
	HCL             = "horuszup/horusx-hcl:v1.1.0"
	Javascript      = "aledj02/horusx-js:v2.10.1"
	Leaks           = "horuszup/horusx-leaks:v1.2.0"
	PHP             = "horuszup/horusx-php:v1.0.1"
	Python          = "aledj02/horusx-python:v2.10.1"
	Ruby            = "horuszup/horusx-ruby:v1.2.0"
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
