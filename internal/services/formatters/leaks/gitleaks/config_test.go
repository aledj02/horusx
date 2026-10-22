package gitleaks

import (
	"strings"
	"testing"
)

func TestCommandUsesPOSIXOutputRedirection(t *testing.T) {
	if strings.Contains(CMD, "&>") {
		t.Fatal("gitleaks command must use POSIX redirection with /bin/sh")
	}
	if !strings.Contains(CMD, "> /tmp/leaks-output.txt 2>&1") {
		t.Fatal("gitleaks command must keep stdout and stderr out of the JSON stream")
	}
}
