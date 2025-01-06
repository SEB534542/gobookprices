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

type Sites []Site

type Site struct {
	Name      string
	Host      string
	Query     string
	Extractor func(doc *goquery.Document) (float64, error)
}

var AllSites = Sites{
	{
		Name:  "Bol.com",
		Host:  HostBol,
		Query: "?searchtext=", // E.g. https://www.bol.com/nl/nl/s/?searchtext=9789021462691
		Extractor: func(doc *goquery.Document) (float64, error) {
			var price string
			// Extract the price from the <meta> tag with itemprop="price"
			doc.Find(`meta[itemprop="price"]`).Each(func(i int, s *goquery.Selection) {
				price, _ = s.Attr("content")
			})

			if price == "" {
				return 0.0, ErrorNotFound
			}

			return strconv.ParseFloat(price, 64)
		}},
	{
		Name:  "Kobo.com",
		Host:  HostKobo,
		Query: "search?query=", // E.g. https://www.kobo.com/nl/en/search?query=9789021462691
		Extractor: func(doc *goquery.Document) (float64, error) {
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
		}},
	{
		Name:  "Onlinebibliotheek.nl",
		Host:  HostOB,
		Query: "zoekresultaten.catalogus.html?q=", // E.g. https://www.onlinebibliotheek.nl/zoekresultaten.catalogus.html?q=9789021462691
		Extractor: func(doc *goquery.Document) (float64, error) {
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
		}},
}

var ErrorNotFound = fmt.Errorf("ISBN not found")

func (xs Sites) GetPrices(isbn string) ([]float64, error) {
	prices := []float64{}
	var err error
	for _, site := range xs {
		price, err1 := site.GetPrice(isbn)
		if err1 != nil {
			err = fmt.Errorf("\n%w\n%w", err, err1)
		}
		prices = append(prices, price)
	}
	return prices, err
}

func (s Site) GetPrice(isbn string) (float64, error) {
	// the mock server for testing does not end with an '/' while the other servers do. So if it does not exist in the host, it is added to prevent errors.
	if s.Host[len(s.Host)-1] != 47 {
		s.Host = fmt.Sprint(s.Host, "/")
	}
	url := fmt.Sprint(s.Host, s.Query, isbn)
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
	return s.Extractor(doc)
}
