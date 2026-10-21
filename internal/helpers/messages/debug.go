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

// Block of messages usage into log of the level debug
const (
	MsgDebugDockerAPIPullNewImage        = "{HORUSX_CLI} Docker pull new image: "
	MsgDebugDockerAPIDownloadWithSuccess = "{HORUSX_CLI} Docker download new image with success: "
	MsgDebugDockerAPIContainerCreated    = "{HORUSX_CLI} Docker create new container: "
	MsgDebugDockerAPIContainerWait       = "{HORUSX_CLI} Docker wait container up..."
	MsgDebugDockerAPIContainerRead       = "{HORUSX_CLI} Docker read container output: "
	MsgDebugDockerAPIFinishedSuccess     = "{HORUSX_CLI} Docker Finished analysis with SUCCESS: "
	MsgDebugDockerAPIFinishedError       = "{HORUSX_CLI} Docker Finished analysis with ERROR: "
	MsgDebugToolStartAnalysis            = "{HORUSX_CLI} Running %s - %s in analysisID: %s"
	MsgDebugToolFinishAnalysis           = "{HORUSX_CLI} %s - %s is finished in analysisID: %s"
	MsgDebugOutputEmpty                  = "{HORUSX_CLI} When format Output it's Empty!"
	MsgDebugConfigFileRunningOnPath      = "{HORUSX_CLI} Config file running on path: "
	MsgDebugFolderOrFileIgnored          = "{HORUSX_CLI} The file or folder was ignored to send analysis:"
	MsgDebugShowConfigs                  = "{HORUSX_CLI} The current configuration for this analysis are:"
	MsgDebugShowWorkdir                  = "{HORUSX_CLI} Using path %s as workdir to run tool %s"
	MsgDebugToolIgnored                  = "{HORUSX_CLI} The tool was ignored for run in this analysis: "
	MsgDebugVulnHashToFix                = "{HORUSX_CLI} Vulnerability Hash expected to be FIXED: "
	MsgDebugDockerImageDoesNotExists     = "{HORUSX_CLI} Image %s does not exists. Pulling from registry"
)
