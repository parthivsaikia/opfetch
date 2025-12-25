package fetcher

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetDistroName(filepath string) (string, error) {
	distroName := ""
	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatalf("Error closing file %v", err)
		}
	}()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "PRETTY_NAME") {

			distroName = strings.Split(line, "=")[1]
		}

	}
	distroName, err = strconv.Unquote(distroName)
	if err != nil {
		return "", err
	}
	return distroName, nil
}
