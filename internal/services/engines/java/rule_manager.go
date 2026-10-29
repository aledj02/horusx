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

package java

import (
	engine "github.com/ZupIT/horusec-engine"

	"github.com/aledj02/horusx/internal/services/engines"
	"github.com/aledj02/horusx/internal/services/engines/jvm"
)

func NewRules() *engines.RuleManager {
	return engines.NewRuleManager(Rules(), extensions())
}

func extensions() []string {
	return []string{".java", ".xml", ".gradle"}
}

// Rules return all rules registred to Java engine.
func Rules() []engine.Rule {
	java := []engine.Rule{
		// And rules
		NewMessageDigestIsCustom(),
		NewInsecureImplementationOfSSL(),
		NewWebViewLoadFilesFromExternalStorage(),
		NewInsecureWebViewImplementation(),
		// NewNoUseSQLCipherAndMatch(),
		// NewNoUseRealmDatabaseWithEncryptionKey(),
		NewNoUseWebviewDebuggingEnable(),
		NewNoListenToClipboard(),
		NewNoCopyContentToClipboard(),
		NewNoUseWebviewIgnoringSSL(),
		NewSQLInjectionWithSqlUtil(),
		NewNoUseFridaServer(),
		// NewNoUseSSLPinningLib(), // SSL pinning is a defensive control, not a vulnerability.
		NewNoUseDexGuardAppDebuggable(),
		NewNoUseDexGuardDebuggerConnected(),
		NewNoUseDexGuardEmulatorDetection(),
		NewNoUseDexGuardWithDebugKey(),
		NewNoUseDexGuardRoot(),
		NewNoUseDexGuard(),
		NewNoUseDexGuardInSigner(),
		NewNoUsePackageWithTamperDetection(),
		NewLoadAndManipulateDexFiles(),
		NewObfuscation(),
		NewExecuteOSCommand(),
		NewTCPServerSocket(),
		NewTCPSocket(),
		NewUDPDatagramPacket(),
		NewUDPDatagramSocket(),
		NewWebViewScriptInterface(),
		NewGetCellInformation(),
		NewGetCellLocation(),
		NewGetSubscriberID(),
		NewGetDeviceID(),
		NewGetSoftwareVersion(),
		NewGetSIMSerialNumber(),
		NewGetSIMProviderDetails(),
		NewGetSIMOperatorName(),
		NewQueryDatabaseOfSMSContacts(),
		NewPotentialPathTraversal(),
		NewJakartaAndPotentialPathTraversal(),
		NewPotentialPathTraversalUsingScalaAPI(),
		NewSMTPHeaderInjection(),
		NewInsecureSMTPSSLConnection(),
		NewAnonymousLDAPBind(),
		NewLDAPEntryPoisoning(),
		NewTrustManagerThatAcceptAnyCertificatesClient(),
		// NewTrustManagerThatAcceptAnyCertificatesServer(),
		// NewTrustManagerThatAcceptAnyCertificatesIssuers(),
		NewXMLParsingVulnerableToXXE(),
		NewIgnoringXMLCommentsInSAML(),
		NewInformationExposureThroughAnErrorMessage(),
		NewHTTPParameterPollution(),
		NewXMLParsingVulnerableToXXEWithDocumentBuilder(),
		NewAWSQueryInjection(),
		NewPotentialTemplateInjectionPebble(),
		NewPotentialTemplateInjectionFreemarker(),
		NewPersistentCookieUsage(),
		NewRequestDispatcherFileDisclosure(),
		NewSpringFileDisclosure(),
		NewStrutsFileDisclosure(),
		NewUnsafeJacksonDeserializationConfiguration(),
		NewObjectDeserializationUsed(),
		NewPotentialCodeScriptInjection(),
		NewPotentialCodeScriptInjectionWithSpringExpression(),
		NewCookieWithoutTheHttpOnlyFlag(),
		NewWebViewWithGeolocationActivated(),
		NewUseOfESAPIEncryptor(),
		NewStaticIV(),
		NewXMLDecoderUsage(),
		NewPotentialXSSInServlet(),
		NewEscapingOfSpecialXMLCharactersIsDisabled(),
		NewDynamicVariableInSpringExpression(),
		NewRSAUsageWithShortKey(),
		NewBlowfishUsageWithShortKey(),
		NewXMLParsingVulnerableToXXEWithXMLInputFactory(),
		NewXMLParsingVulnerableToXXEWithSAXParserFactory(),
		NewXMLParsingVulnerableToXXEWithTransformerFactory(),
		// NewXMLParsingVulnerableToXXEWithSchemaFactory(),
		NewXMLParsingVulnerableToXXEWithDom4j(),
		NewXMLParsingVulnerableToXXEWithJdom2(),
		NewClassesShouldNotBeLoadedDynamically(),
		// NewHostnameVerifierVerifyShouldNotAlwaysReturnTrue(),
		NewXPathExpressionsShouldNotBeVulnerableToInjectionAttacks(),
		NewServerHostnamesShouldBeVerifiedDuringSSLTLSConnections(),
		NewServerHostnamesShouldBeVerifiedDuringSSLTLSConnectionsWithSimpleEmail(),
		NewFunctionCallsShouldNotBeVulnerableToPathInjectionAttacks(),
		NewServerHostnamesShouldBeVerifiedDuringSSLTLSConnectionsWithMail(),
		// NewServerHostnamesShouldBeVerifiedDuringSSLTLSConnectionsWithJakartaMail(),
		NewHTTPResponseHeadersShouldNotBeVulnerableToInjectionAttacks(),
		NewLDAPAuthenticatedAnalyzeYourCode(),
		NewSecureRandomSeedsShouldNotBePredictable(),
		NewExceptionsShouldNotBeThrownFromServletMethods(),
		NewActiveMQConnectionFactoryVulnerableToMaliciousCodeDeserialization(),
		NewOpenSAML2ShouldBeConfiguredToPreventAuthenticationBypass(),
		NewHttpServletRequestGetRequestedSessionIdShouldNotBeUsed(),
		NewJakartaAndHttpServletRequestGetRequestedSessionIdShouldNotBeUsed(),
		// NewWebApplicationsShouldHotHaveAMainMethod(), // Spring Boot applications legitimately have main.
		// NewJakartaAndWebApplicationsShouldHotHaveAMainMethod(),

		// Or Rules
		NewFileIsWorldReadable(),
		NewFileIsWorldWritable(),
		// NewNoWriteExternalContent(), // Storage API usage alone does not prove exposure.
		NewNoUseIVsWeak(),
		// NewRootDetectionCapabilities(), // Root detection is a defensive capability.
		NewJARURLConnection(),
		// NewSetOrReadClipboardData(),
		// NewMessageDigest(),
		NewOverlyPermissiveFilePermission(),
		NewCipherGetInstanceInsecure(),
		NewVulnerableRemoteCodeExecutionSpringFramework(),

		// Regular rules
		// NewHiddenElements(), // Hidden UI elements are normal application behavior.
		NewWeakCypherBlockMode(),
		NewPossibleFileWithVulnerabilityWhenOpen(),
		NewWeakHash(),
		// NewSensitiveInformationNotEncrypted(), // MODE_PRIVATE does not imply unencrypted data.
		NewInsecureRandomNumberGenerator(),
		NewNoDefaultHash(),
		// NewLayoutParamsFlagSecure(), // FLAG_SECURE is a defensive control.
		// NewNoUseSQLCipher(), // getWritableDatabase does not prove missing encryption.
		// NewPreventTapJackingAttacks(), // This setting is itself defensive.
		// NewPreventWriteSensitiveInformationInTmpFile(), // Temp-file creation alone is benign.
		// NewGetWindowFlagSecure(), // FLAG_SECURE is a defensive control.
		// NewLoadingNativeCode(), // JNI loading is normal and not unsafe by itself.
		NewDynamicClassAndDexloading(),
		// NewCryptoImport(), // Importing crypto libraries is not a vulnerability.
		// NewStartingService(), // Service APIs require permission/export analysis.
		// NewSendingBroadcast(), // Broadcast APIs require permission/data-flow analysis.
		// NewLocalFileOperations(), // Local I/O APIs are normal application behavior.
		// NewInterProcessCommunication(), // IPC requires permission/export analysis.
		NewURLRewritingMethod(),
		NewOverlyPermissiveCORSPolicy(),
		NewHostnameVerifierThatAcceptAnySignedCertificates(),
		// NewDefaultHttpClient(), // The old pattern matched the recommended client.
		NewWeakSSLContext(),
		NewSQLInjection(),
		NewDisablingHTMLEscaping(),
		NewSQLInjectionWithTurbine(),
		NewSQLInjectionWithHibernate(),
		NewSQLInjectionWithJDO(),
		NewSQLInjectionWithJPA(),
		NewSQLInjectionWithSpringJDBC(),
		NewSQLInjectionWithJDBC(),
		NewLDAPInjection(),
		NewUnsafeHashEquals(),
		NewPotentialExternalControl(),
		NewBadHexadecimalConcatenation(),
		NewNullCipherInsecure(),
		NewUnvalidatedRedirect(),
		NewRequestMappingMethodsNotPublic(),
		NewLDAPDeserializationNotDisabled(),
		NewDatabasesPasswordNotProtected(),
		NewVulnerableRemoteCodeInjectionApacheLog4j(),
		NewUncheckedClassInstatiation(),
	}
	return append(java, jvm.Rules()...)
}
