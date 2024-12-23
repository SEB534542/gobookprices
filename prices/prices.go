package gobookprices

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

const HostBol = "https://www.bol.com/nl/nl/s/"

func getPriceBol(hostUrl, isbn string) (float64, error) {
	query := "?searchtext="
	url := fmt.Sprint(hostUrl, query, isbn) // E.g. https://www.bol.com/nl/nl/s/?searchtext=9789021462691

	resp, err := http.Get(url)
	if err != nil {
		return 0.0, fmt.Errorf("getting price from %s failed: %w", url, err)
	}
	defer resp.Body.Close()

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return 0.0, fmt.Errorf("failed to load HTML from '%s': %w", url, err)
	}

	// Initialize a variable to store the price as float64
	var price float64

	// Extract the price from the <meta> tag with itemprop="price"
	doc.Find(`meta[itemprop="price"]`).Each(func(i int, s *goquery.Selection) {
		content, exists := s.Attr("content")
		if exists {
			// Convert the string value to float64
			parsedPrice, err := strconv.ParseFloat(content, 64)
			if err == nil {
				price = parsedPrice
			}
		}
	})

	return price, nil
}
