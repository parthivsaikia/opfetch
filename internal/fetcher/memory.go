package fetcher

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type MemoryInfo struct {
	TotalMem float64
	UsedMem  float64
	FreeMem  float64
}

func GetMemory(filepath string) (*MemoryInfo, error) {
	var totalMem int
	var freeMem int
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Fatalf("Error while closing file %v", err)
		}
	}()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		l := scanner.Text()
		if strings.HasPrefix(l, "MemTotal") {
			totalMem = parseLine(l)
		}
		if strings.HasPrefix(l, "MemAvailable") {
			freeMem = parseLine(l)

		}
	}

	usedMem := totalMem - freeMem

	return &MemoryInfo{
		TotalMem: convertKbtoGb(totalMem),
		UsedMem:  convertKbtoGb(usedMem),
		FreeMem:  convertKbtoGb(freeMem),
	}, nil

}

func parseLine(l string) int {
	parts := strings.Fields(l)
	if len(parts) < 2 {
		return 0
	}
	memoryInKb, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0
	}
	return memoryInKb
}

func convertKbtoGb(kb int) float64 {
	memInGb := float64(kb) / (1024 * 1024)
	return math.Round(memInGb*100) / 100
}
