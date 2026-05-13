package output

import (
	"encoding/json"
	"os"

	"github.com/gilbertonyenwezi2026-oss/go-wikipedia-focused-crawler/internal/model"
)

func WriteJSONLines(filename string, items []model.PageItem) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)

	for _, item := range items {
		if err := encoder.Encode(item); err != nil {
			return err
		}
	}

	return nil
}
