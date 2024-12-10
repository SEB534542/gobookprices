package gobookprices

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const goodreadsURL = "https://www.goodreads.com/review/list/"

type book struct {
	title       string
	author      string
	isbn        string // Remove not relevant?
	workUrl     string // This is the data-resource-id from goodreads, used as identifier for a work.
	editionsUrl string // This is the url to all editions of the work
}

type edition struct {
	isbn     string
	format   string
	language string
}

// getBooksFromPage takes an url to a specific page in a goodreads public account and returns the books and an error.
// example url: fmt.Sprintf("https://www.goodreads.com/review/list/%s?page=1&per_page=20", yourGoodreadsId)
func getBooksFromPage(url string) ([]book, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("getting books from %s failed: %w", url, err)
	}
	defer resp.Body.Close()

	// Read the body.
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("creating document failed: %w", err)
	}

	nocontent := strings.TrimSpace(doc.Find(".greyText.nocontent.stacked").Text())
	if nocontent == "No matching items!" {
		// Reset the page to break the loop.
	}

	books := []book{}

	doc.Find("#booksBody tr").Each(func(j int, l *goquery.Selection) {
		title := strings.ReplaceAll(strings.TrimSuffix(strings.TrimSpace(l.Find(".title .value").Text()), "\n        *"), "\n        ", " ")
		author := strings.TrimSuffix(strings.TrimSpace(l.Find(".author .value").Text()), "\n        *")
		isbn := strings.TrimSpace(l.Find(".isbn .value").Text())
		if len(isbn) < 1 {
			// Try the ISBN 13.
			isbn = strings.TrimSpace(l.Find(".isbn13 .value").Text())
			if len(isbn) < 1 {
				// Try the asin.
				isbn = strings.TrimSpace(l.Find(".asin .value").Text())
			}
		}
		workUrl, _ := l.Find(".js-tooltipTrigger").Attr("data-resource-id")
		workUrl = fmt.Sprintf("/book/show/%s", workUrl)
		editionUrl, _ := l.Find("td.field.format a").Attr("href")

		books = append(books, book{
			title:       title,
			author:      author,
			isbn:        isbn,
			workUrl:     workUrl,
			editionsUrl: editionUrl,
		})
	})
	return books, nil
}

// getBooks takes a goodreadsID and a specific shelf. It returns the books from that shelf and an error.
// Example url https://www.goodreads.com/review/list/68156753?page=1&per_page=20&shelf=to-read
func getBooks(goodreadsId, shelf string) ([]book, error) {
	books := []book{}
	for p := 1; ; p++ {
		url := fmt.Sprintf("%s%s?page=%d&per_page=20&shelf=%s", goodreadsURL, goodreadsId, p, shelf)
		b, err := getBooksFromPage(url)
		if len(b) == 0 {
			break
		}
		books = append(books, b...)
		if err != nil {
			return books, err
		}
		time.Sleep(time.Second) // wait one second, to limit queries to goodreads (not sure if this is necessary)
	}
	return books, nil
}

// getEditionsFromPage takes an url containing editions on goodreads, scapes and returns the editions and an error.
// Example url: fmt.Sprintf("https://www.goodreads.com/work/editions/%s?page=1&per_page=10", editionsUrl)
func getEditionsFromPage(url string) ([]edition, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("getting books from %s failed: %v", url, err)
	}
	defer resp.Body.Close()

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return []edition{}, fmt.Errorf("failed to get editions from '%s' failed: %w", url, err)
	}

	var editions []edition

	// Regex to clean up the ISBN, removing anything within parentheses
	reISBN := regexp.MustCompile(`\s*\(.*?\)`)

	// Find each "elementList clearFix"
	doc.Find(".elementList.clearFix").Each(func(i int, selection *goquery.Selection) {
		var e edition

		// Extract format (e.g., "Hardcover, 396 pages")
		selection.Find(".dataRow").Each(func(j int, dataRow *goquery.Selection) {
			text := strings.TrimSpace(dataRow.Text())
			if strings.Contains(text, "pages") {
				// Take the part before the comma
				e.format = strings.Split(text, ",")[0]
			}
		})

		// Extract ISBN
		selection.Find(".dataTitle:contains('ISBN')").Each(func(j int, dataTitle *goquery.Selection) {
			dataValue := dataTitle.SiblingsFiltered(".dataValue").Text()
			cleanedISBN := strings.TrimSpace(reISBN.ReplaceAllString(dataValue, ""))
			e.isbn = cleanedISBN
		})

		// Extract language (Edition language)
		selection.Find(".dataTitle:contains('Edition language')").Each(func(j int, dataTitle *goquery.Selection) {
			dataValue := dataTitle.SiblingsFiltered(".dataValue").Text()
			e.language = strings.TrimSpace(dataValue)
		})

		// Add to editions slice if at least one field is populated
		if e.format != "" || e.isbn != "" || e.language != "" {
			editions = append(editions, e)
		}
	})
	return editions, nil
}
