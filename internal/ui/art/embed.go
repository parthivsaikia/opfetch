package art

import (
	"embed"
	"fmt"
)

//go:embed *.txt
var artFS embed.FS

func GetCharacterArt(name string) string {
	filename := fmt.Sprintf("%s.txt", name)
	content, err := artFS.ReadFile(filename)
	if err != nil {
		return "Character not found"
	}
	return string(content)
}

func GetWantedText() string {
	content, err := artFS.ReadFile("wanted.txt")
	if err != nil {
		return "Wanted text not found"
	}
	return string(content)
}

func GetDeadText() string {
	content, err := artFS.ReadFile("dead.txt")
	if err != nil {
		return "Dead text not found"
	}
	return string(content)
}
