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

package analysis_test

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/ZupIT/horusec-devkit/pkg/enums/languages"
	"github.com/ZupIT/horusec-devkit/pkg/utils/logger"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/aledj02/horusx/config"
	"github.com/aledj02/horusx/e2e/analysis"
	customimages "github.com/aledj02/horusx/internal/entities/custom_images"
	"github.com/aledj02/horusx/internal/utils/testutil"
)

const (
	isWindows        = runtime.GOOS == "windows"
	isDarwin         = runtime.GOOS == "darwin"
	horusxConfigName = "horusx-config.json"
)

var _ = Describe("Run a complete horusx analysis when build tools locally", func() {
	tmpDir := CreateHorusxConfigAndReturnTMPDirectory()

	Describe("Running e2e tests when build tools locally", func() {
		allTestsCases := analysis.NewTestCase()

		for idx, testCase := range allTestsCases {
			if !isDarwin {
				horusxConfigFilePathBuildLocally := filepath.Join(tmpDir, horusxConfigName)
				RunTestCase(testCase, horusxConfigFilePathBuildLocally, "(build locally)", idx, len(allTestsCases))

				time.Sleep(2 * time.Second)
			}

			horusxConfigFilePathDownloadFromDockerHub := filepath.Join(allTestsCases[idx].Command.Flags[testutil.StartFlagProjectPath], horusxConfigName)
			RunTestCase(testCase, horusxConfigFilePathDownloadFromDockerHub, "(download from dockerhub)", idx, len(allTestsCases))

			time.Sleep(2 * time.Second)
		}
	})

	AfterEach(func() {
		if err := os.RemoveAll(tmpDir); err != nil {
			Fail(fmt.Sprintf("Error on remove tmp folder: %v", err))
		}
	})
})

func RunTestCase(testCase *analysis.TestCase, horusxConfigFilePath, runType string, idx, lenTestCase int) {
	testCase.Command.Flags[testutil.GlobalFlagConfigFilePath] = horusxConfigFilePath
	if isWindows && testCase.RequiredDocker {
		logger.LogInfo("Tool ignored because is required docker: ", runType, testCase.Tool, fmt.Sprintf("[%v/%v]", idx+1, lenTestCase))
		return
	}
	logger.LogInfo("Preparing test for run e2e: ", runType, testCase.Tool, fmt.Sprintf("[%v/%v]", idx+1, lenTestCase))
	session, err := testCase.RunAnalysisTestCase()
	if err != nil {
		Fail(fmt.Sprintf("Error on run analysis Test Case %s: %v", runType, err))
	}
	session.Wait(testutil.AverageTimeoutAnalyzeForExamplesFolder)
	testCase.Command.Output = string(session.Out.Contents())
	testCase.Command.ExitCode = session.ExitCode()

	It(fmt.Sprintf("Execute command without error for %s: %s", runType, testCase.Tool), func() {
		Expect(testCase.Command.ExitCode).Should(Equal(0))
	})

	It(fmt.Sprintf("Validate is outputs expected exists on tool %s: %s", runType, testCase.Tool), func() {
		for _, outputExpected := range testCase.Expected.OutputsContains {
			Expect(testCase.Command.Output).Should(
				ContainSubstring(outputExpected),
				fmt.Sprintf("The output [%s] not exist in output", outputExpected),
				testCase.Command.Output,
			)
		}
	})

	It(fmt.Sprintf("Validate is outputs not expected exists on tool %s: %s", runType, testCase.Tool), func() {
		for _, outputNotExpected := range testCase.Expected.OutputsNotContains {
			Expect(testCase.Command.Output).ShouldNot(
				ContainSubstring(outputNotExpected),
				fmt.Sprintf("The output [%s] exist in output", outputNotExpected),
				testCase.Command.Output,
			)
		}
	})
}

func CreateHorusxConfigAndReturnTMPDirectory() string {
	tmpDir := filepath.Join(os.TempDir(), "horusx-analysis-e2e-"+uuid.NewString())
	horusxConfigPath := filepath.Join(tmpDir, horusxConfigName)

	if err := os.MkdirAll(tmpDir, os.ModePerm); err != nil {
		Fail(fmt.Sprintf("Error on create tmp folder: %v", err))
	}

	horusxConfigFile, err := os.Create(horusxConfigPath)
	if err != nil {
		Fail(fmt.Sprintf("Error on create config file for scan on e2e test %v", err))
	}

	horusxConfig := config.New()

	newCustomImages := customimages.Default()
	for k := range newCustomImages {
		if k == languages.CSharp {
			newCustomImages[k] = strings.ToLower("local-csharp:local")
		} else {
			newCustomImages[k] = strings.ToLower(fmt.Sprintf("local-%s:local", k))
		}
	}

	horusxConfig.CustomImages = newCustomImages
	horusxConfigBytes, err := json.MarshalIndent(horusxConfig.ToMapLowerCase(), "", "  ")
	if err != nil {
		Fail(fmt.Sprintf("Error on marshall horusx-config.json %v", err))
	}

	_, err = horusxConfigFile.Write(horusxConfigBytes)
	if err != nil {
		Fail(fmt.Sprintf("Error on write config file for scan on e2e test %v", err))
	}
	return tmpDir
}
