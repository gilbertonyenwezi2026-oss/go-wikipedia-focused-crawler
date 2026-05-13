# Go Wikipedia Focused Crawler

## Project Overview

This project implements a Go-based focused web crawler and scraper using the [Colly](https://github.com/gocolly/colly) framework. The crawler collects text from selected Wikipedia pages related to intelligent systems, robotics, artificial intelligence, software agents, robotic process automation, and related research topics.

The program retrieves web pages from a predefined list of Wikipedia URLs, extracts readable article text, removes HTML markup, and writes the results to a JSON lines file named `items.jl`.

This project was developed for MSDS 431 as a Go-based alternative to a Python/Scrapy crawler.

---

## Management Problem

A technology firm is building an internal online library and knowledge base focused on its current research and development efforts in intelligent systems and robotics. Some of the content for this knowledge base will be collected from the World Wide Web.

The firm previously created a Python/Scrapy crawler that collects pages sequentially. However, this approach can become slow when the number of target pages grows from a small test list to hundreds or thousands of web pages.

Management wants to evaluate whether Go can provide a faster and more scalable solution. Go is well suited for this use case because it supports lightweight concurrency through goroutines and channels. Web crawling and scraping involve many independent network requests, making them a good candidate for concurrent processing.

This project demonstrates how Go and Colly can be used to build a focused crawler that retrieves Wikipedia pages, extracts text, and stores the output in a format suitable for later search indexing, document storage, knowledge graph development, or text analytics.

---

## Assignment Objective

The objective of this assignment is to create a Go-based web crawler/scraper that produces results similar to the Python/Scrapy demonstration.

The program must:

1. Use a predefined list of Wikipedia URLs related to intelligent systems and robotics.
2. Retrieve each web page from the World Wide Web.
3. Extract readable article text from each page.
4. Ignore images, videos, audio, navigation menus, and HTML markup.
5. Store the scraped results in a JSON lines file.
6. Include tests for critical program components.
7. Provide documentation for running, testing, and building the application.
8. Include an executable application file.
9. Document any use of generative AI tools.

---

## Technologies Used

- Go
- Colly web scraping framework
- JSON Lines / NDJSON output format
- Go testing package
- Git and GitHub

---

## Data Sources

The crawler uses the following Wikipedia pages as the initial focused crawl list:

```go
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
```

These pages represent the firm’s research focus on robotics, intelligent systems, artificial intelligence applications, and automation.

---

## Repository Structure

```text
go-wikipedia-focused-crawler/
│
├── README.md
├── go.mod
├── go.sum
├── items.jl
├── crawler.log
├── gen_ai_log.txt
│
├── cmd/
│   └── crawler/
│       └── main.go
│
├── internal/
│   ├── model/
│   │   └── page.go
│   │
│   ├── scraper/
│   │   ├── scraper.go
│   │   └── scraper_test.go
│   │
│   └── output/
│       ├── jsonlines.go
│       └── jsonlines_test.go
│
├── testdata/
│   └── sample_wikipedia.html
│
└── bin/
    └── crawler.exe
```

### File and Folder Descriptions

| File or Folder | Description |
|---|---|
| `cmd/crawler/main.go` | Main application entry point. Defines the URL list, configures Colly, crawls pages, extracts text, and writes output. |
| `internal/model/page.go` | Defines the `PageItem` struct used for scraped page records. |
| `internal/scraper/scraper.go` | Contains helper functions for cleaning text and counting words. |
| `internal/scraper/scraper_test.go` | Unit tests for text cleaning and word counting. |
| `internal/output/jsonlines.go` | Writes scraped page records to a JSON lines file. |
| `internal/output/jsonlines_test.go` | Unit tests for JSON lines output. |
| `items.jl` | Final JSON lines output file containing scraped Wikipedia page text. |
| `crawler.log` | Runtime log file showing visited pages, scraped page titles, word counts, errors, and elapsed time. |
| `gen_ai_log.txt` | Documentation of generative AI tool usage. |
| `bin/crawler.exe` | Windows executable build of the crawler application. |

---

## How the Crawler Works

The program uses the Colly framework to visit each Wikipedia URL.

The main steps are:

1. Create a Colly collector.
2. Restrict crawling to `en.wikipedia.org`.
3. Enable asynchronous crawling.
4. Limit parallel requests to avoid overwhelming the target website.
5. Visit each URL in the predefined list.
6. Extract the page title from the `h1` element.
7. Extract article paragraphs from the main Wikipedia content area.
8. Clean whitespace from the extracted text.
9. Count words for each scraped page.
10. Store each page as a `PageItem`.
11. Write all page records to `items.jl`.
12. Write runtime details to `crawler.log`.

---

## Output Format

The main deliverable is:

```text
items.jl
```

This is a JSON lines file. Each line contains one JSON object representing one scraped web page.

Example format:

```json
{"url":"https://en.wikipedia.org/wiki/Robotics","title":"Robotics","text":"Robotics is the interdisciplinary study and practice...","word_count":3500}
{"url":"https://en.wikipedia.org/wiki/Robot","title":"Robot","text":"A robot is a machine...","word_count":4200}
```

Each JSON object includes:

| Field | Description |
|---|---|
| `url` | Source Wikipedia URL. |
| `title` | Page title extracted from the Wikipedia article. |
| `text` | Cleaned article text extracted from the page. |
| `word_count` | Number of words in the extracted text. |

The file uses the `.jl` extension because the assignment requires a JSON lines file.

---

## Why JSON Lines?

JSON lines, also known as NDJSON or newline-delimited JSON, is useful for document processing because each line is a complete record. This format is commonly used for:

- Document databases
- Search engines
- Knowledge bases
- Text analytics
- Machine learning pipelines
- Information retrieval systems
- Knowledge graph development

Unlike a single large JSON array, JSON lines files can be processed one record at a time.

---

## Setup Instructions

### Prerequisites

Make sure Go is installed.

Check your Go version:

```bash
go version
```

This project was built using Go modules.

---

## Install Dependencies

From the project root folder, run:

```bash
go mod tidy
```

If Colly has not already been installed, run:

```bash
go get github.com/gocolly/colly/v2
go mod tidy
```

---

## How to Run the Program

From the project root folder, run:

```bash
go run ./cmd/crawler
```

Expected console output:

```text
Crawler completed successfully
Pages scraped: 10
Output file: items.jl
Log file: crawler.log
```

After the program runs, confirm that these files were created or updated:

```text
items.jl
crawler.log
```

---

## How to Run Tests

Run all unit tests from the project root folder:

```bash
go test ./...
```

Expected output should show passing tests for the scraper and output packages.

Example:

```text
ok      go-wikipedia-focused-crawler/internal/scraper
ok      go-wikipedia-focused-crawler/internal/output
```

The tests cover:

- Text cleaning
- Word counting
- JSON lines file writing

---

## How to Build the Executable

For Windows:

```bash
go build -o bin/crawler.exe ./cmd/crawler
```

Run the executable:

```bash
bin\crawler.exe
```

For macOS or Linux:

```bash
go build -o bin/crawler ./cmd/crawler
```

Run the executable:

```bash
./bin/crawler
```

---

## Logging

The program writes runtime information to:

```text
crawler.log
```

The log file includes:

- Program start message
- URLs visited
- Failed requests, if any
- Page titles scraped
- Word counts
- Number of pages scraped
- Total elapsed runtime

Example log entries:

```text
Starting Go Wikipedia focused crawler
Visiting: https://en.wikipedia.org/wiki/Robotics
Scraped page: Robotics | Words: 3500
Crawler completed successfully
Pages scraped: 10
Elapsed time: 3.52s
```

---

## Performance Notes

The Go crawler uses Colly with asynchronous crawling enabled:

```go
colly.Async(true)
```

The program also uses a domain limit rule to control parallelism:

```go
c.Limit(&colly.LimitRule{
    DomainGlob:  "*wikipedia.org*",
    Parallelism: 4,
})
```

This means multiple Wikipedia pages can be requested concurrently while still behaving responsibly toward the target website.

For a small list of 10 Wikipedia pages, total runtime may be dominated by network latency rather than CPU processing. However, the Go design is more scalable than a purely sequential crawler because it can keep multiple requests active at the same time.

For larger lists of hundreds or thousands of URLs, the concurrent Go/Colly implementation should provide stronger performance advantages over a sequential approach.

---

## Python Scrapy vs. Go Colly Observations

The original Python/Scrapy demonstration uses a list of Wikipedia URLs, retrieves each page, extracts web content, and writes results to a JSON lines file. This Go version follows the same general workflow but uses the Colly framework instead of Scrapy.

Key differences:

| Area | Python Scrapy | Go Colly |
|---|---|---|
| Language | Python | Go |
| Framework | Scrapy | Colly |
| Concurrency model | Scrapy engine and asynchronous networking | Goroutines and Colly async crawling |
| Output | JSON lines file | JSON lines file |
| Primary use case | Web crawling and scraping | Fast, lightweight web crawling and scraping |
| Assignment value | Original prototype | Go-based replacement implementation |

The Go version is designed to be simple, readable, and scalable while producing the required `items.jl` file.

---

## Design Notes

The project separates responsibilities into small modules:

- The `model` package defines the structure of scraped records.
- The `scraper` package handles text cleaning and word counting.
- The `output` package handles JSON lines file creation.
- The `main` package controls crawling, scraping, logging, and runtime measurement.

This design improves readability, testing, and maintainability.

---

## Testing and Software Metrics

This project includes unit tests for critical components.

Tested components include:

1. `CleanText()`  
   Ensures extra spaces, line breaks, and tabs are normalized.

2. `CountWords()`  
   Ensures word counts are calculated correctly.

3. `WriteJSONLines()`  
   Ensures the program can write valid JSON lines output.

The program also generates basic runtime metrics:

- Number of pages scraped
- Word count for each page
- Total elapsed time

These metrics are recorded in `crawler.log`.

---

## Application Behavior

The application does not require user input. It runs using the predefined list of Wikipedia URLs.

The program should:

1. Start successfully.
2. Visit each Wikipedia URL.
3. Extract article title and text.
4. Write output to `items.jl`.
5. Write logs to `crawler.log`.
6. Exit cleanly.

If a page request fails, the error is logged and the program continues processing the remaining URLs.

---

## Limitations

This project is intentionally focused on Wikipedia pages and uses Wikipedia-specific CSS selectors.

Current limitations include:

- The crawler is limited to the predefined URL list.
- It does not recursively follow links to new Wikipedia pages.
- It ignores images, tables, videos, and references.
- It extracts primarily paragraph text.
- It is not designed for websites outside Wikipedia without selector adjustments.

These limitations are appropriate for the assignment because the requirement is to retrieve text information from a focused list of Wikipedia pages.

---

## Future Improvements

Possible future improvements include:

1. Reading URLs from a configuration file.
2. Adding command-line arguments for output file names.
3. Adding recursive crawling with depth limits.
4. Adding duplicate URL detection.
5. Adding retry logic for failed requests.
6. Adding richer metadata such as scrape timestamp, categories, and page length.
7. Adding CSV export in addition to JSON lines.
8. Adding a comparison benchmark against the Python/Scrapy version.
9. Adding text preprocessing for downstream NLP or knowledge graph work.
10. Adding support for additional domains beyond Wikipedia.

---

## Gen AI Tools

Generative AI tools were used during the development of this assignment.

ChatGPT was used to support the following tasks:

- Explaining how Colly can be used for focused web crawling
- Drafting starter code for the crawler
- Creating text cleaning and JSON lines output functions
- Suggesting unit tests
- Debugging Go module and PowerShell command issues
- Preparing a short Gen AI tools disclosure

All AI-generated suggestions were reviewed, edited, and tested before being included in the final project.


---

## References

Colly. “Elegant Scraper and Crawler Framework for Golang.” GitHub.  
https://github.com/gocolly/colly

Hoare, C. A. R. *Communicating Sequential Processes*. Prentice Hall International, 1985.

McConnell, Steve. *Code Complete: A Practical Handbook of Software Construction*. 2nd ed. Microsoft Press, 2004.

Mitchell, Ryan. *Web Scraping with Python: Collecting More Data from the Modern Web*. 2nd ed. O’Reilly Media, 2018.

The Go Authors. “The Go Programming Language Documentation.”  
https://go.dev/doc/

Wikipedia. “Robotics.”  
https://en.wikipedia.org/wiki/Robotics

Wikipedia. “Robot.”  
https://en.wikipedia.org/wiki/Robot

Wikipedia. “Reinforcement Learning.”  
https://en.wikipedia.org/wiki/Reinforcement_learning

Wikipedia. “Robot Operating System.”  
https://en.wikipedia.org/wiki/Robot_Operating_System

Wikipedia. “Intelligent Agent.”  
https://en.wikipedia.org/wiki/Intelligent_agent

Wikipedia. “Software Agent.”  
https://en.wikipedia.org/wiki/Software_agent

Wikipedia. “Robotic Process Automation.”  
https://en.wikipedia.org/wiki/Robotic_process_automation

Wikipedia. “Chatbot.”  
https://en.wikipedia.org/wiki/Chatbot

Wikipedia. “Applications of Artificial Intelligence.”  
https://en.wikipedia.org/wiki/Applications_of_artificial_intelligence

Wikipedia. “Android Robot.”  
https://en.wikipedia.org/wiki/Android_(robot)

---

## Submission Notes

The GitHub repository should include:

- Source code
- Go module files
- Unit tests
- README documentation
- `items.jl` output file
- `crawler.log`
- Executable file
- Gen AI tools documentation

The cloneable GitHub repository URL submitted in Canvas should end with `.git`.

Example:

```text
https://github.com/gilbertonyenwezi2026-oss/go-wikipedia-focused-crawler.git
```
