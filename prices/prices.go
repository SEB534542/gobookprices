package gobookprices

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	HostBol  = "https://www.bol.com/nl/nl/s/"
	HostKobo = "https://www.kobo.com/nl/en/"
	HostOB   = "https://www.onlinebibliotheek.nl/"
)

var ErrorNotFound = fmt.Errorf("ISBN not found")

// getPriceBol takes a URL to a Host and the isbn you want to get the price from Bol from. It returns the price and an error.
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
		} else {
			err = fmt.Errorf("failed to find price in html from '%s'", url)
		}
	})

	return price, nil
}

func getPriceKobo(hostUrl, isbn string) (float64, error) {
	query := "search?query="
	url := fmt.Sprint(hostUrl, "/", query, isbn) // E.g. https://www.kobo.com/nl/en/search?query=9789021462691

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

	// Find the price from the 'data-price' attribute
	price := ""
	doc.Find("div.price-format").EachWithBreak(func(i int, s *goquery.Selection) bool {
		price, _ = s.Attr("data-price")
		if price != "" {
			return false // Stop after finding the first price
		}
		return true
	})

	if price == "" {
		return 0.0, ErrorNotFound
	}
	// Clean and convert the price to float64
	price = strings.ReplaceAll(price, "€", "")  // Remove €
	price = strings.ReplaceAll(price, ",", ".") // Replace ',' with '.'
	priceFloat, err := strconv.ParseFloat(strings.TrimSpace(price), 64)
	if err != nil {
		err = fmt.Errorf("error converting price from '%s' to float: %w", url, err)
	}

	return priceFloat, err
}

// getPriceOB takes a URL to a Host and the isbn you want to get the price from Bol from. It returns the price and an error.
func getPriceOB(hostUrl, isbn string) (float64, error) {
	query := "zoekresultaten.catalogus.html?q="
	url := fmt.Sprint(hostUrl, query, isbn) // E.g. // https://www.onlinebibliotheek.nl/zoekresultaten.catalogus.html?q=9789021462691

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

	// Check if the div with class 'content list-big' has any content
	div := doc.Find("div.content.list-big")
	if div.Length() == 0 {
		// div not found
		return 0.0, ErrorNotFound
	}

	// Check if it contains any child elements or text
	content := strings.TrimSpace(div.Text())
	if content == "" {
		// div is empty
		return 0.0, ErrorNotFound
	}

	return 0.0, nil
}
