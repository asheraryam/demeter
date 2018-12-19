package lib

import (
	"fmt"
	"time"
)

// SearchResult contains the results from a calibre query
type SearchResult struct {
	TotalNum  int    `json:"total_num"`
	Offset    int    `json:"offset"`
	Sort      string `json:"sort"`
	LibraryID string `json:"library_id"`
	SortOrder string `json:"sort_order"`
	Vl        string `json:"vl"`
	BaseURL   string `json:"base_url"`
	Num       int    `json:"num"`
	BookIds   []int  `json:"book_ids"`
	Query     string `json:"query"`
}

// CalibreBook is a book from calibre
type CalibreBook struct {
	UUID          string            `json:"uuid"`
	Title         string            `json:"title"`
	ApplicationID int               `json:"application_id"`
	TitleSort     string            `json:"title_sort"`
	Cover         string            `json:"cover"`
	Pubdate       string            `json:"pubdate"`
	MainFormat    map[string]string `json:"main_format"`
	AuthorSort    string            `json:"author_sort"`
	Authors       []string          `json:"authors"`
	Timestamp     time.Time         `json:"timestamp"`
	Languages     []string          `json:"languages"`
	LastModified  time.Time         `json:"last_modified"`
	Thumbnail     string            `json:"thumbnail"`
	Formats       []string          `json:"formats"`
}

// BooksQueryResult is
type BooksQueryResult map[string]CalibreBook

// Host describes all attributes related to a host
type Host struct {
	ID            int    `storm:"id,increment"`
	URL           string `storm:"unique"`
	Downloads     int
	Scrapes       int
	LastScrape    time.Time
	LastDownload  time.Time
	Added         time.Time
	ScrapeResults []ScrapeResult
	Active        bool
}

// ScrapeResult is the result of a single scrape attempt
type ScrapeResult struct {
	Start     time.Time
	End       time.Time
	Success   bool
	Results   int
	Downloads int
}

// Print prints a scrapeResult in a nicely formatted way
func (s *ScrapeResult) Print() {
	niceDuration := s.End.Sub(s.Start).String()
	fmt.Printf(` - Started:   %s
   Duration:  %s
   Success:   %t
   Downloads: %d
   Results:   %d`, s.Start.Format(time.RFC3339), niceDuration, s.Success, s.Downloads, s.Results)
	fmt.Println()
}

// Book is a oversimplified representation of a book
type Book struct {
	ID       int `storm:"id,increment"`
	Added    time.Time
	Hash     string `storm:"unique"`
	SourceID int    `storm:"index"`
	Author   string
	Title    string
}

// Print prints a host in a nicely formatted way
func (h *Host) Print(verbose bool) {
	fmt.Printf(`ID:          %d
URL:         %s
Scrapes:     %d
Downloads:   %d
Active:      %t`, h.ID, h.URL, h.Scrapes, h.Downloads, h.Active)
	fmt.Println()
	if verbose {
		fmt.Println("Scrape results: ")
		if h.Scrapes == 0 || len(h.ScrapeResults) == 0 {
			fmt.Println(" - none")
		} else {
			for _, lastScrape := range h.ScrapeResults {
				lastScrape.Print()
			}
		}
	}
}
