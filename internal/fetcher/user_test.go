package fetcher_test

import (
	"os"
	"testing"

	"github.com/parthivsaikia/opfetch/internal/fetcher"
)

func TestGetUsername(t *testing.T) {
	expectedUser := os.Getenv("USER")
	got, err := fetcher.GetUsername()
	if err != nil {
		t.Fatal(err)
	}
	if expectedUser != got {
		t.Errorf("Expected username to %s but got %s", expectedUser, got)
	}
}
