package gobookprices

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type ScraperConfig struct {
	SearchURL      string
	TitleSelector  string
	PriceSelector  string
	LinkSelector   string
}

type EditionInfo struct {
	Title string
	Price string
	Link  string
}

func ScrapeData(config ScraperConfig, isbn string) ([]EditionInfo, error) {
	// Replace placeholder with ISBN
	url := fmt.Sprintf(config.SearchURL, isbn)

	// Fetch the page
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch URL: %w", err)
	}
	defer resp.Body.Close()

	// Handle HTTP errors
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received HTTP status: %d", resp.StatusCode)
	}

	// Parse HTML
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTML: %w", err)
	}

	// Extract data
	var editions []EditionInfo
	doc.Find(config.TitleSelector).Each(func(i int, s *goquery.Selection) {
		title := s.Text()
		price := s.Parent().Find(config.PriceSelector).Text()
		link, _ := s.Attr("href")

		editions = append(editions, EditionInfo{
			Title: title,
			Price: price,
			Link:  link,
		})
	})

	return editions, nil
}
