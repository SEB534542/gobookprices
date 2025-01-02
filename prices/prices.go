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

func getPrice(hostUrl, query, isbn string, extractor func(doc *goquery.Document) (float64, error)) (float64, error) {
	// the mock server for testing does not end with an '/' while the other servers do. So if it does not exist in the host, it is added to prevent errors.
	if hostUrl[len(hostUrl)-1] != 47 {
		hostUrl = fmt.Sprint(hostUrl, "/")
	}
	url := fmt.Sprint(hostUrl, query, isbn)
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

	// Extract the price from the HTML document and return output
	return extractor(doc)
}

// getPriceBol takes a URL to a Host and the isbn you want to get the price from Bol from. It returns the price and an error.
func getPriceBol(hostUrl, isbn string) (float64, error) {
	query := "?searchtext=" // E.g. https://www.bol.com/nl/nl/s/?searchtext=9789021462691
	extractor := func(doc *goquery.Document) (float64, error) {
		var price string
		// Extract the price from the <meta> tag with itemprop="price"
		doc.Find(`meta[itemprop="price"]`).Each(func(i int, s *goquery.Selection) {
			price, _ = s.Attr("content")
		})

		if price == "" {
			return 0.0, ErrorNotFound
		}

		return strconv.ParseFloat(price, 64)
	}

	return getPrice(hostUrl, query, isbn, extractor)
}

func getPriceKobo(hostUrl, isbn string) (float64, error) {
	query := "search?query=" // E.g. https://www.kobo.com/nl/en/search?query=9789021462691
	extractor := func(doc *goquery.Document) (float64, error) {
		// Find the price from the 'data-price' attribute
		price := ""
		doc.Find("div.price-format").EachWithBreak(func(i int, s *goquery.Selection) bool {
			price, _ = s.Attr("data-price")
			return price == ""
		})

		if price == "" {
			return 0.0, ErrorNotFound
		}
		// Clean and convert the price to float64
		price = strings.ReplaceAll(price, "€", "")  // Remove €
		price = strings.ReplaceAll(price, ",", ".") // Replace ',' with '.'
		return strconv.ParseFloat(strings.TrimSpace(price), 64)
	}
	return getPrice(hostUrl, query, isbn, extractor)
}

// getPriceOB takes a URL to a Host and the isbn you want to get the price from Bol from. It returns the price and an error.
func getPriceOB(hostUrl, isbn string) (float64, error) {
	query := "zoekresultaten.catalogus.html?q=" // E.g. https://www.onlinebibliotheek.nl/zoekresultaten.catalogus.html?q=9789021462691
	extractor := func(doc *goquery.Document) (float64, error) {
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

		return 0.0, nil // this source has no prices, so zero is returned
	}
	return getPrice(hostUrl, query, isbn, extractor)
}
