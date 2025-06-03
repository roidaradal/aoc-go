package aoc15

import (
	"fmt"
	"regexp"
	"strings"

	. "github.com/roidaradal/aoc-go/aoc"
)

func Day08() {
	words := data08(true)

	// Part 1
	total := 0
	for _, word := range words {
		total += len(word) - len(parseString(word))
	}
	fmt.Println(total)

	// Part 2
	total = 0
	for _, word := range words {
		total += len(expandString(word)) - len(word)
	}
	fmt.Println(total)
}

func data08(full bool) []string {
	return ReadLines(full)
}

func parseString(text string) string {
	hex := regexp.MustCompile(`\\x[0-9a-f]{2}`)
	text = text[1 : len(text)-1]
	text = hex.ReplaceAllString(text, ".")
	text = strings.ReplaceAll(text, `\"`, `"`)
	text = strings.ReplaceAll(text, `\\`, ".")
	return text
}

func expandString(text string) string {
	chars := make([]string, 0)
	for _, x := range text {
		if x == '"' {
			chars = append(chars, `\"`)
		} else if x == '\\' {
			chars = append(chars, `\\`)
		} else {
			chars = append(chars, string(x))
		}
	}
	return fmt.Sprintf(`"%s"`, strings.Join(chars, ""))
}
