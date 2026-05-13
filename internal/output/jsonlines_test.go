package output

import (
	"bufio"
	"os"
	"testing"

	"github.com/gilbertonyenwezi2026-oss/go-wikipedia-focused-crawler/internal/model"
)

func TestWriteJSONLines(t *testing.T) {
	filename := "test_items.jl"

	items := []model.PageItem{
		{
			URL:       "https://example.com",
			Title:     "Example",
			Text:      "This is example text.",
			WordCount: 4,
		},
	}

	err := WriteJSONLines(filename, items)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	defer os.Remove(filename)

	file, err := os.Open(filename)
	if err != nil {
		t.Fatalf("could not open output file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0

	for scanner.Scan() {
		lineCount++
	}

	if lineCount != 1 {
		t.Errorf("expected 1 JSON line, got %d", lineCount)
	}
}
