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

package java

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	engine "github.com/ZupIT/horusec-engine"
	"github.com/ZupIT/horusec-engine/text"
	"github.com/stretchr/testify/require"
)

func TestContextualRulesAvoidCommonFalsePositives(t *testing.T) {
	tests := []struct {
		name     string
		rule     engine.Rule
		source   string
		findings int
	}{
		{
			name: "random used for UI does not report",
			rule: NewInsecureRandomNumberGenerator(),
			source: `class Example {
				void shuffle() {
					java.util.Random random = new java.util.Random();
				}
			}`,
		},
		{
			name: "random used for token reports",
			rule: NewInsecureRandomNumberGenerator(),
			source: `class Example {
				String createToken() {
					java.util.Random random = new java.util.Random();
					return String.valueOf(random.nextLong());
				}
			}`,
			findings: 1,
		},
		{
			name: "hashCode used by collections does not report",
			rule: NewNoDefaultHash(),
			source: `class Example {
				int bucket(Object value) {
					return value.hashCode();
				}
			}`,
		},
		{
			name: "hashCode used for password reports",
			rule: NewNoDefaultHash(),
			source: `class Example {
				int passwordHash(String password) {
					return password.hashCode();
				}
			}`,
			findings: 1,
		},
		{
			name: "custom hostname verifier validates the host",
			rule: NewHostnameVerifierThatAcceptAnySignedCertificates(),
			source: `class Example implements HostnameVerifier {
				public boolean verify(String host, SSLSession session) {
					return session.getPeerHost().equals(host);
				}
			}`,
		},
		{
			name: "hostname verifier accepting every host reports",
			rule: NewHostnameVerifierThatAcceptAnySignedCertificates(),
			source: `class Example implements HostnameVerifier {
				public boolean verify(String host, SSLSession session) {
					return true;
				}
			}`,
			findings: 1,
		},
		{
			name: "unrelated private method does not report",
			rule: NewRequestMappingMethodsNotPublic(),
			source: `class Example {
				@RequestMapping("/public")
				public void endpoint() {}

				private void helper() {}
			}`,
		},
		{
			name: "private request mapping reports",
			rule: NewRequestMappingMethodsNotPublic(),
			source: `class Example {
				@RequestMapping("/private")
				private void endpoint() {}
			}`,
			findings: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			filename := filepath.Join(t.TempDir(), "Example.java")
			require.NoError(t, os.WriteFile(filename, []byte(test.source), 0o644))

			findings, err := engine.NewEngine(0, "*").Run(context.Background(), filename, test.rule)
			require.NoError(t, err)
			require.Len(t, findings, test.findings)
		})
	}
}

func TestDeprecatedNoisyRulesAreNotRegistered(t *testing.T) {
	deprecatedIDs := map[string]bool{
		"HS-JAVA-28":  true,
		"HS-JAVA-96":  true,
		"HS-JAVA-97":  true,
		"HS-JAVA-101": true,
		"HS-JAVA-103": true,
		"HS-JAVA-109": true,
		"HS-JAVA-113": true,
		"HS-JAVA-116": true,
		"HS-JAVA-117": true,
		"HS-JAVA-118": true,
		"HS-JAVA-119": true,
		"HS-JAVA-120": true,
		"HS-JAVA-121": true,
		"HS-JAVA-123": true,
		"HS-JAVA-124": true,
		"HS-JAVA-125": true,
		"HS-JAVA-126": true,
		"HS-JAVA-127": true,
		"HS-JAVA-128": true,
	}

	for _, rule := range Rules() {
		textRule, ok := rule.(*text.Rule)
		require.True(t, ok, "expected Java rule to be a text rule")
		require.False(t, deprecatedIDs[textRule.Metadata.ID], "deprecated noisy rule %s is still registered", textRule.Metadata.ID)
	}
}
