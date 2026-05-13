package model

type PageItem struct {
	URL       string `json:"url"`
	Title     string `json:"title"`
	Text      string `json:"text"`
	WordCount int    `json:"word_count"`
}
