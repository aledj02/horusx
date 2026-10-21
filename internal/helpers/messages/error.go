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

package messages

// Block of messages usage into error response
//
// nolint:lll
const (
	MsgErrorPathNotValid                        = "invalid path:"
	MsgErrorJSONOutputFilePathNotValidExtension = "Output File path not valid file of type:"
	MsgErrorJSONOutputFilePathNotValidUnknown   = "Output File path is required or is invalid:"
	MSgErrorGitLeaksPermissionDenied            = "error: could not lock config file .git/config: Permission denied"
	MsgErrorSeverityNotValid                    = "Type of severity not valid. See severities enable:"
	MsgErrorAskForUserCancelled                 = "{HORUSX_CLI} Operation was canceled by user"
	MsgVulnerabilityTypeToShowInvalid           = "{HORUSX_CLI} Error on validate vulnerability type is wrong type: "
	MsgErrorRunToolInDocker                     = "{HORUSX_CLI} Error to execute tool %s | analysisID -> %s | output -> %s"
	MsgErrorInvalidWorkDir                      = "{HORUSX_CLI} Workdir is nil! Check the configuration and try again"
	MsgErrorParseStringToToolsConfig            = "{HORUSX_CLI} Error when try parse tools config string to entity. Returning default values"
	MsgErrorNotFoundRequirementsTxt             = "file requirements.txt was not found in your Python project. It would be a good idea to commit it so Horusx can check for vulnerabilities"
	MsgErrorPackageLockJSONNotFound             = "file package-lock.json was not found in your Javascript project. If you use NPM to handle your dependencies, it would be a good idea to commit it so Horusx can check for vulnerabilities"
	MsgErrorYarnLockNotFound                    = "file yarn.lock file was not found in your Javascript project. If you use Yarn to handle your dependencies, it would be a good idea to commit it so Horusx can check for vulnerabilities"
	MsgErrorYarnProcess                         = "{HORUSX_CLI} Yarn returned an error: "
	MsgErrorGetFilenameByExt                    = "Could not get filename by extension: "
	MsgErrorNancyRateLimit                      = `{HORUSX_CLI} Nancy tool failed to query the GitHub API for updates.
This is most likely due to GitHub rate-limiting on unauthenticated requests.
To make authenticated requests please:
  1. Generate a token at https://github.com/settings/tokens
  2. Set the token by setting the GITHUB_TOKEN environment variable.
Instructions for generating a token can be found at:
https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line. `
)

// Block of messages usage into log of the level error
//
// nolint:lll
const (
	MsgErrorFalsePositiveNotValid       = "False positive is not valid because is duplicated in risk accept:"
	MsgErrorRiskAcceptNotValid          = "Risk Accept is not valid because is duplicated in false positive:"
	MsgErrorWhenCheckRequirementsGit    = "{HORUSX_CLI} Error when check if git requirement it's ok!"
	MsgErrorWhenCheckRequirementsDocker = "{HORUSX_CLI} Error when check if docker requirement it's ok!"
	MsgErrorWhenCheckDockerRunning      = "{HORUSX_CLI} Error when check if docker is running."
	MsgErrorWhenDockerIsLowerVersion    = "{HORUSX_CLI} Your docker version is below of: "
	MsgErrorWhenGitIsLowerVersion       = "{HORUSX_CLI} Your git version is below of: "
	MsgErrorRemoveAnalysisFolder        = "{HORUSX_CLI} Error when remove analysis project inside .horusx"
	MsgErrorDetectLanguage              = "{HORUSX_CLI} Error when detect language"
	MsgErrorCopyProjectToHorusxAnalysis = "{HORUSX_CLI} Error when copy project to .horusx folder"
	MsgErrorGenerateJSONFile            = "{HORUSX_CLI} Error when try parse horusx analysis to output"
	MsgErrorDockerPullImage             = "{HORUSX_CLI} Error when pull new image: "
	MsgErrorDockerListImages            = "{HORUSX_CLI} Error when list all images enable: "
	MsgErrorDockerCreateContainer       = "{HORUSX_CLI} Error when create container of analysis: "
	MsgErrorDockerStartContainer        = "{HORUSX_CLI} Error when start container of analysis: "
	MsgErrorDockerListAllContainers     = "{HORUSX_CLI} Error when list all containers of analysis: "
	MsgErrorDockerRemoveContainer       = "{HORUSX_CLI} Error when remove container of analysis: "
	MsgErrorGitCommitAuthorsExecute     = "{HORUSX_CLI} Error when execute commit author command: "
	MsgErrorGitCommitAuthorsParseOutput = "{HORUSX_CLI} Error when to parse output to commit author struct: "
	MsgErrorParseStringToWorkDir        = "{HORUSX_CLI} Error when try parse workdir string to entity." +
		"Returning default values"
	MsgErrorDeferFileClose                   = "{HORUSX_CLI} Error defer file close: "
	MsgErrorSetHeadersOnConfig               = "{HORUSX-CLI} Error on set headers on configurations"
	MsgErrorReplayWrong                      = "{HORUSX-CLI} Error on set reply, Please type Y or N. Your current response was: "
	MsgErrorErrorOnCreateConfigFile          = "{HORUSX-CLI} Error on create config file: "
	MsgErrorErrorOnReadConfigFile            = "{HORUSX-CLI} Error on read config file on path: "
	MsgErrorFailedToPullImage                = "{HORUSX_CLI} Failed to pull docker image"
	MsgErrorWhileParsingCustomImages         = "{HORUSX_CLI} Error when parsing custom images config. Using default values"
	MsgErrorSettingLogFile                   = "{HORUSX_CLI} Error when setting log file"
	MsgErrorGetRelativePathFromFile          = "{HORUSX_CLI} Error when get relative path of file"
	MsgErrorGetDependencyCodeFilepathAndLine = "{HORUSX_CLI} Error when get dependency code filepath and line"
	MsgErrorGetDependencyInfo                = "{HORUSX_CLI} Error when get dependency code info"
	MsgErrorBundlerNotAccessDB               = "{HORUSX_CLI} BundlerAudit cannot access database in github: "
)
