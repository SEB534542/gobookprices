package gobookprices

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const goodreadsURL = "https://www.goodreads.com/review/list/"

type book struct {
	id     string // this is the data-resource-id from goodreads, used as identifier for a book.
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
		id, _ := l.Find(".js-tooltipTrigger").Attr("data-resource-id")
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
			id:     id,
			title:  title,
			author: author,
			isbn:   isbn,
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
		fmt.Printf("getting data from %s\n", url)
		b, err := getBooksFromPage(url)
		if len(b) == 0 {
			break
		}
		fmt.Printf("%v books found\n", len(b))
		books = append(books, b...)
		if err != nil {
			return books, err
		}
		time.Sleep(time.Second) // wait one second, to limit queries to goodreads (not sure if this is necessary)
	}
	return books, nil
}
