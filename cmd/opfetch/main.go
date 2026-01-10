package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"

	"github.com/parthivsaikia/opfetch/internal/bounty"
	"github.com/parthivsaikia/opfetch/internal/fetcher"
	"github.com/parthivsaikia/opfetch/internal/ui"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	versionFlag := flag.Bool("version", false, "Display the version of opfetch")
	flag.Parse()
	if *versionFlag {
		fmt.Printf("opfetch %s (commit: %s, built at: %s)\n", version, commit, date)
		return
	}
	logger := slog.New(slog.NewJSONHandler(os.Stderr, nil))

	username, err := fetcher.GetUsername()
	if err != nil {
		username = "Monkey D. Luffy"
		logger.Warn("Failed to fetch username", slog.Any("error", err))
	}

	bounty := bounty.CalculateBounty()

	hostname, err := fetcher.GetHostname()
	if err != nil {
		hostname = "Going Merry"
		logger.Warn("Failed to fetch hostname", slog.Any("error", err))
	}

	uptime, err := fetcher.GetUptime("/proc/uptime")
	if err != nil {
		uptime.Days = 0
		uptime.Hours = 0
		uptime.Minutes = 0
		uptime.Seconds = 0
		logger.Warn("Failed to fetch uptime", slog.Any("error", err))
	}

	packages, err := fetcher.GetPackageCount()
	if err != nil {
		packages = 0
		logger.Warn("Failed to fetch package count", slog.Any("error", err))
	}

	distro, err := fetcher.GetDistroName("/etc/os-release")
	if err != nil {
		distro = "Straw Hat"
		logger.Warn("Failed to fetch distroname", slog.Any("error", err))
	}

	shell := fetcher.GetShell()

	ui.DrawUI(username, bounty, hostname, uptime.String(), packages, distro, shell)
}
