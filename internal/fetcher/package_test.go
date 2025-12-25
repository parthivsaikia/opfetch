package fetcher

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGetPackage(t *testing.T) {
	t.Run("Arch linux", func(t *testing.T) {
		originalPath := pacmanPath

		dir, err := os.MkdirTemp("", "mockpacman")
		if err != nil {
			t.Fatal(err)
		}

		defer os.RemoveAll(dir)

		pacmanPath = dir

		defer func() {
			pacmanPath = originalPath
		}()

		packages := []string{"vi", "firefox", "curl", "code"}

		for _, p := range packages {
			err := os.Mkdir(filepath.Join(dir, p), 0755)
			if err != nil {
				t.Fatal(err)
			}
		}

		expectedPackageCount := 4

		gotPackageCount, err := GetPackageCount()

		if err != nil {
			t.Fatal(err)
		}

		if expectedPackageCount != gotPackageCount {
			t.Errorf("Expected package count %d but got %d", expectedPackageCount, gotPackageCount)
		}
	})

	t.Run("Ubuntu", func(t *testing.T) {
		originalPath := dpkgPath

		defer func() {
			dpkgPath = originalPath
		}()

		file, err := os.CreateTemp("", "mockdpkg")
		if err != nil {
			t.Fatal(err)
		}

		dpkgPath = file.Name()

		dpkgData := `

Package: bash
Status: install ok installed
Priority: required
Section: shells
Installed-Size: 7164
Maintainer: Ubuntu Developers <ubuntu-devel-discuss@lists.ubuntu.com>
Architecture: amd64
Version: 5.1-6ubuntu1
Depends: libc6 (>= 2.34), libtinfo6 (>= 6)
Description: GNU Bourne Again SHell
 Bash is an sh-compatible command language interpreter.

Package: curl
Status: install ok installed
Priority: optional
Section: web
Installed-Size: 448
Maintainer: Ubuntu Developers <ubuntu-devel-discuss@lists.ubuntu.com>
Architecture: amd64
Version: 7.81.0-1ubuntu1
Depends: libc6 (>= 2.34), libcurl4
Description: command line tool for transferring data with URL syntax
 Curl is a command line tool for transferring data with URL syntax.

Package: htop
Status: install ok installed
Priority: optional
Section: utils
Installed-Size: 320
Maintainer: Ubuntu Developers <ubuntu-devel-discuss@lists.ubuntu.com>
Architecture: amd64
Version: 3.2.1-1ubuntu1
Depends: libc6 (>= 2.34), libncursesw6
Description: interactive processes viewer
 Htop is an interactive text-mode process viewer.
		`

		if _, err := file.WriteString(dpkgData); err != nil {
			t.Fatal(err)
		}

		file.Close()

		originalPacmanPath := pacmanPath
		pacmanPath = "/nothing"

		defer func() {
			pacmanPath = originalPacmanPath
		}()

		expectedPackageCount := 3
		gotPackageCount, err := GetPackageCount()
		if err != nil {
			t.Fatal(err)
		}
		if expectedPackageCount != gotPackageCount {
			t.Errorf("Expected package count %d but got %d", expectedPackageCount, gotPackageCount)
		}
	})
}
