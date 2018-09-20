package main

import (
	"strings"
	"testing"
)

func TestOutputVersion(t *testing.T) {
	if want, got := Version, OutputVersion(); !strings.Contains(got, want) {
		t.Fatalf("Enexpected version string returned want: %s, got: %s", want, got)
	}
}
