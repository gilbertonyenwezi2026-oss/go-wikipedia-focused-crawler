package main

import (
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/gocolly/colly/v2"

	"github.com/gilbertonyenwezi2026-oss/go-wikipedia-focused-crawler/internal/model"
	"github.com/gilbertonyenwezi2026-oss/go-wikipedia-focused-crawler/internal/output"
	"github.com/gilbertonyenwezi2026-oss/go-wikipedia-focused-crawler/internal/scraper"
)

func main() {
	start := time.Now()

	logFile, err := os.Create("crawler.log")
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)
	log.Println("Starting Go Wikipedia focused crawler")

	urls := []string{
		"https://en.wikipedia.org/wiki/Robotics",
		"https://en.wikipedia.org/wiki/Robot",
		"https://en.wikipedia.org/wiki/Reinforcement_learning",
		"https://en.wikipedia.org/wiki/Robot_Operating_System",
		"https://en.wikipedia.org/wiki/Intelligent_agent",
		"https://en.wikipedia.org/wiki/Software_agent",
		"https://en.wikipedia.org/wiki/Robotic_process_automation",
		"https://en.wikipedia.org/wiki/Chatbot",
		"https://en.wikipedia.org/wiki/Applications_of_artificial_intelligence",
		"https://en.wikipedia.org/wiki/Android_(robot)",
	}

	var items []model.PageItem
	var mu sync.Mutex

	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"),
		colly.Async(true),
	)

	err = c.Limit(&colly.LimitRule{
		DomainGlob:  "*wikipedia.org*",
		Parallelism: 4,
	})
	if err != nil {
		log.Fatal(err)
	}

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting:", r.URL.String())
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request failed:", r.Request.URL.String(), err)
	})

	c.OnHTML("html", func(e *colly.HTMLElement) {
		url := e.Request.URL.String()

		title := scraper.CleanText(e.ChildText("h1"))

		var paragraphs []string
		e.ForEach("div.mw-parser-output > p", func(_ int, el *colly.HTMLElement) {
			text := scraper.CleanText(el.Text)
			if text != "" {
				paragraphs = append(paragraphs, text)
			}
		})

		fullText := scraper.CleanText(strings.Join(paragraphs, " "))
		wordCount := scraper.CountWords(fullText)

		item := model.PageItem{
			URL:       url,
			Title:     title,
			Text:      fullText,
			WordCount: wordCount,
		}

		mu.Lock()
		items = append(items, item)
		mu.Unlock()

		log.Printf("Scraped page: %s | Words: %d\n", title, wordCount)
	})

	for _, url := range urls {
		if err := c.Visit(url); err != nil {
			log.Println("Visit error:", url, err)
		}
	}

	c.Wait()

	if err := output.WriteJSONLines("items.jl", items); err != nil {
		log.Fatal(err)
	}

	elapsed := time.Since(start)

	log.Println("Crawler completed successfully")
	log.Println("Pages scraped:", len(items))
	log.Println("Elapsed time:", elapsed)

	println("Crawler completed successfully")
	println("Pages scraped:", len(items))
	println("Output file: items.jl")
	println("Log file: crawler.log")
	println("Elapsed time:", elapsed.String())
}
