package gobookprices

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type book struct {
	title  string
	author string
	isbn   string
}

// getBooksFromPage takes an url to a specific page in a goodreads public account and returns the books and an error.
// example url: fmt.Sprintf("https://www.goodreads.com/review/list/%s?page=1&per_page=20", yourGoodreadsId)
func getBooksFromPage(url string) ([]book, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("getting books from %s failed: %v", url, err)
	}
	defer resp.Body.Close()

	// Read the body.
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("creating document failed: %v", err)
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
		books = append(books, book{
			title:  title,
			author: author,
			isbn:   isbn,
		})
	})
	return books, nil
}

// func getBooks(goodreadsId string) ([]book, error){

// }