package tfsec

import (
	"strings"
	"testing"
)

func TestCommandWritesJSONToFile(t *testing.T) {
	if !strings.Contains(CMD, "--out=/tmp/tfsec-result-ANALYSISID.json") {
		t.Fatal("tfsec command must write JSON to a file")
	}
	if !strings.Contains(CMD, "cat /tmp/tfsec-result-ANALYSISID.json") {
		t.Fatal("tfsec command must emit the JSON file after execution")
	}
}
