package fetcher

import (
	"bufio"
	"log"
	"os"
	"strings"
)

var (
	pacmanPath = "/var/lib/pacman/local"
	dpkgPath   = "/var/lib/dpkg/status"
)

func GetPackageCount() (int, error) {
	if isExist(pacmanPath) {
		return getPacmanPackageCount()
	}

	if isExist(dpkgPath) {
		return getDpkgPackageCount()
	}

	return 0, nil

}

func isExist(filepath string) bool {
	_, err := os.Stat(filepath)
	return err == nil
}

func getPacmanPackageCount() (int, error) {
	count := 0
	entries, err := os.ReadDir(pacmanPath)
	if err != nil {
		if os.IsNotExist(err) {
			return 0, nil
		}
		return 0, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			count++
		}
	}
	return count, nil

}

func getDpkgPackageCount() (int, error) {
	count := 0
	file, err := os.Open(dpkgPath)
	if err != nil {
		return 0, err
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatalf("Error while closing file %v", err)
		}
	}()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "Package") {
			count++
		}
	}
	return count, nil
}
