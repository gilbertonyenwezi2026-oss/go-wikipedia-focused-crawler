package scraper

import (
	"regexp"
	"strings"
)

func CleanText(text string) string {
	space := regexp.MustCompile(`\s+`)
	cleaned := space.ReplaceAllString(text, " ")
	return strings.TrimSpace(cleaned)
}

func CountWords(text string) int {
	fields := strings.Fields(text)
	return len(fields)
}
