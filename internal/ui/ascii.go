package ui

import (
	"github.com/common-nighthawk/go-figure"
)

func generateAscii(text string, asciiFont string) string {
	asciiArt := figure.NewFigure(text, asciiFont, true)
	return asciiArt.ColorString()
}
