package fetcher_test

import (
	"os/exec"
	"strings"
	"testing"

	"github.com/parthivsaikia/opfetch/internal/fetcher"
)

func TestGetHostname(t *testing.T) {
	cmd := exec.Command("hostname")
	output, err := (cmd.Output())
	expectedHost := strings.TrimSpace(string(output))
	if err != nil {
		t.Fatal(err)
	}
	gotHostname, err := fetcher.GetHostname()
	if err != nil {
		t.Fatal(err)
	}
	if string(expectedHost) != gotHostname {
		t.Errorf("Expected hostname %s but got %s", expectedHost, gotHostname)
	}
}
