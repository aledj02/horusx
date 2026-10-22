package scs_test

import (
	"strings"
	"testing"

	"github.com/ZupIT/horusec/internal/services/formatters/csharp/scs"
)

func TestCommandUsesPOSIXOutputRedirection(t *testing.T) {
	if strings.Contains(scs.CMD, "&>") {
		t.Fatal("security-scan command must use POSIX redirection with /bin/sh")
	}
	if !strings.Contains(scs.CMD, "> /tmp/scs-run-output-ANALYSISID.txt 2>&1") {
		t.Fatal("security-scan command must keep diagnostics out of the JSON stream")
	}
}
