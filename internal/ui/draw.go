package ui

import (
	"fmt"
)

func DrawUI(username string, bounty float64, hostname string, uptime string, packages int, distro string, shell string) {
	fmt.Println("WANTED")
	fmt.Printf("%s@%s\n", username, hostname)
	fmt.Printf("$ %d\n", int(bounty))
	fmt.Printf("Uptime: %s\n", uptime)
	fmt.Printf("Treasure: %d\n", packages)
	fmt.Printf("Distro: %s\n", distro)
	fmt.Printf("Shell: %s\n", shell)
}
