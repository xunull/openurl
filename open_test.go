package openurl

import (
	"testing"
)

func TestOpenUrl(t *testing.T) {
	err, stdout := Open("https://www.google.com")
	if err != nil {
		t.Logf("failed: %v", err)
		t.Logf("stdout: %s", stdout)
	}
}
