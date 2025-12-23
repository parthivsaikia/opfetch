package fetcher

import (
	"os"
)

func GetPackageCount(filepath string) (int, error) {
	pacmanFilePath := "/var/lib/pacman/local"

}

func isExist(filepath string) bool {
	_, err := os.Stat(filepath)
	return err == nil
}

func getPacmanPackage() (int, error) {
	count := 0
	entries, err := os.ReadDir("/var/lib/pacman/local")
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
