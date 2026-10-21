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

package goengine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRules_NoBuiltins(t *testing.T) {
	mgr := NewRules()
	assert.NotNil(t, mgr)
	assert.Empty(t, mgr.GetAllRules(),
		"Go strategy is LIGERO: 0 built-in rules. Custom rules are loaded at runtime via --custom-rules-path.")
}

func TestExtensions(t *testing.T) {
	assert.Equal(t, []string{".go"}, NewRules().GetAllExtensions())
}
