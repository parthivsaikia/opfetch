package ui

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/parthivsaikia/opfetch/internal/ui/art"
	"golang.org/x/term"
)

var (
	paperColor  = lipgloss.Color("#A5907C")
	textColor   = lipgloss.Color("#31231B")
	labelColor  = lipgloss.Color("#5C4033") // Darker brown for labels
	accentColor = lipgloss.Color("#8B6F47") // Medium brown for accents
)

func DrawUI(username string, bounty float64, hostname string, uptime string, packages int, distro string, shell string) {
	wantedHeader := art.GetWantedText()
	artStr := art.GetCharacterArt("luffy")

	// Calculate the widest part of your poster
	contentWidth := max(lipgloss.Width(wantedHeader), lipgloss.Width(artStr))
	posterWidth := contentWidth + 6 // Add some breathing room

	// 2. STYLES
	// Base style for every line in the poster
	sectionStyle := lipgloss.NewStyle().
		Foreground(textColor).
		Background(paperColor).
		Align(lipgloss.Center).
		Width(posterWidth)

	// 3. COMPONENTS
	// Header
	header := sectionStyle.
		Bold(true).
		Render(wantedHeader)

	// Image (Luffy)
	image := sectionStyle.
		PaddingTop(1).
		PaddingBottom(1).
		Render(artStr)

	// Dead or Alive text
	deadText := art.GetDeadText()
	deadSection := sectionStyle.
		Bold(true).
		Render(deadText)

	// Name
	name := username + "@" + hostname
	nameText := sectionStyle.
		Bold(true).
		Render(name)

	// Bounty (above the separator)
	bountyStr := fmt.Sprintf("฿ %.0f", bounty)
	bountyText := sectionStyle.
		Bold(true).
		PaddingTop(1).
		Render(bountyStr)

	// Info separator line
	separatorLine := strings.Repeat("─", posterWidth-6)
	separator := sectionStyle.
		Foreground(accentColor).
		PaddingTop(1).
		PaddingBottom(1).
		Render(separatorLine)

	// Create table for system info
	cellStyle := lipgloss.NewStyle().
		Foreground(textColor).
		Background(paperColor).
		Padding(0)

	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(paperColor).Background(paperColor)).
		StyleFunc(func(row, col int) lipgloss.Style {
			return cellStyle
		}).
		Rows(
			[]string{"OS", ":", distro},
			[]string{"Kernel", ":", "Linux"},
			[]string{"Uptime", ":", uptime},
			[]string{"Packages", ":", fmt.Sprintf("%d", packages)},
			[]string{"Shell", ":", shell},
		)

	// Render table and center it
	tableStr := t.String()
	infoSection := sectionStyle.
		PaddingTop(1).
		PaddingBottom(1).
		Render(tableStr)

	// 4. ASSEMBLY
	poster := lipgloss.JoinVertical(lipgloss.Center,
		header,
		image,
		deadSection,
		nameText,
		bountyText,
		separator,
		infoSection,
	)

	// 5. RENDER
	// Center the poster in the terminal window
	termWidth, termHeight, _ := term.GetSize(int(os.Stdout.Fd()))
	fmt.Println(lipgloss.Place(termWidth, termHeight, lipgloss.Center, lipgloss.Center, poster))
}

// Simple max helper function (Go 1.21+ has this built-in)
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
