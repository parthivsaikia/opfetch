package bounty

import (
	"log/slog"
	"os"
	"strings"

	"github.com/parthivsaikia/opfetch/internal/fetcher"
)

func CalculateBounty() float64 {
	logger := slog.New(slog.NewJSONHandler(os.Stderr, nil))

	var ram float64
	memoryInfo, err := fetcher.GetMemory("/proc/meminfo")
	if err != nil {
		logger.Error("Error fetching memory", "error", err)
		ram = 0 // Fallback to 0 if failed
	} else {
		ram = memoryInfo.TotalMem
	}

	packageCount, err := fetcher.GetPackageCount()
	if err != nil {
		logger.Error("Error fetching packages", "error", err)
		packageCount = 0
	}

	distroname, err := fetcher.GetDistroName("/etc/os-release")
	if err != nil {
		logger.Error("Error fetching distro", "error", err)
		distroname = "unknown"
	}

	shell := fetcher.GetShell()

	return bountyFormula(ram, float64(packageCount), distroname, shell)
}

func bountyFormula(ram float64, packagesCount float64, distroname string, shell string) float64 {
	distroMultiplier := 1.0
	shellMultiplier := 1.0

	name := strings.ToLower(distroname)

	if strings.Contains(name, "arch") || strings.Contains(name, "gentoo") {
		distroMultiplier = 1.5
	} else if strings.Contains(name, "debian") || strings.Contains(name, "fedora") || strings.Contains(name, "ubuntu") {
		distroMultiplier = 1.2
	}

	shell = strings.ToLower(shell)
	if strings.Contains(shell, "zsh") || strings.Contains(shell, "fish") {
		shellMultiplier = 1.1
	}

	base := 50000000.0 // 50 Million base
	ramValue := ram * 20000.0
	pkgValue := packagesCount * 10000.0

	return (base + ramValue + pkgValue) * distroMultiplier * shellMultiplier
}
