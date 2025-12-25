package fetcher

import (
	"os"
	"strings"
)

func GetShell() string {
	shellStr := os.Getenv("SHELL")
	strs := strings.Split(shellStr, "/")
	return strs[len(strs)-1]
}
