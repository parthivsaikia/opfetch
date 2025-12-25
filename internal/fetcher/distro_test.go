package fetcher_test

import (
	"os"
	"testing"

	"github.com/parthivsaikia/opfetch/internal/fetcher"
)

func TestGetDistroName(t *testing.T) {

	t.Run("get distro name in test file", func(t *testing.T) {
		os_release_data := `
NAME="One Piece OS"
PRETTY_NAME="One Piece OS"
ID=arch
BUILD_ID=rolling
ANSI_COLOR="38;2;23;147;209"
HOME_URL="https://archlinux.org/"
DOCUMENTATION_URL="https://wiki.archlinux.org/"
SUPPORT_URL="https://bbs.archlinux.org/"
BUG_REPORT_URL="https://gitlab.archlinux.org/groups/archlinux/-/issues"
PRIVACY_POLICY_URL="https://terms.archlinux.org/docs/privacy-policy/"
LOGO=archlinux-logo
		`
		testFileName := "test-os-release"
		file, err := os.CreateTemp("", testFileName)
		if err != nil {
			t.Fatal(err)
		}

		defer func() {
			if err := os.Remove(file.Name()); err != nil {
				t.Fatalf("Error while removing file %v", err)
			}
		}()

		if _, err := file.Write([]byte(os_release_data)); err != nil {
			t.Fatal(err)
		}

		if err := file.Close(); err != nil {
			t.Fatal(err)
		}

		distroName := "One Piece OS"
		got, err := fetcher.GetDistroName(file.Name())
		if err != nil {
			t.Fatal(err)
		}

		if distroName != got {
			t.Errorf("Expected distro name %s but got %s", distroName, got)
		}

	})
}
