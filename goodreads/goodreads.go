package gobookprices

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const HostUrl = "https://www.goodreads.com/"

type Library struct {
	hostUrl   string
	ListUrl   string
	Shelf     string
	Formats   []string
	Languages []string
	Books     []Book
}

// Book represents a book on a goodread list. It can be a work or an edition.
type Book struct {
	Title       string    // Title of book.
	Author      string    // Author of book.
	Isbn        string    // Isbn of this book. Note: remove as only edition isbn is relevant?
	WorkUrl     string    // This is the url to the work on goodreads.
	EditionsUrl string    // This is the url to all editions of the work.
	Editions    []Edition // This contains all editions that belong to the work.
}

// Edition represents an edition of a work.
type Edition struct {
	Isbn     string
	Format   string
	Language string
}

var EbookFormats = []string{"ebook", "Kindle Edition"} // EbooksFormats contains all formats that represent a form a ebook on goodreads.

// getBooksFromPage takes an url to a specific page in a goodreads public account and returns the books and an error.
// example url: "https://www.goodreads.com/review/list/68156753?page=1&per_page=20&shelf=to-read"
func getBooksFromPage(url string) ([]Book, error) {
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

	books := []Book{}

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
		// editionUrl, _ := l.Find("td.field.format a").Attr("href") // commented-out because this only works if you've accepted the cookies.

		books = append(books, Book{
			Title:   title,
			Author:  author,
			Isbn:    isbn,
			WorkUrl: workUrl,
			//	EditionsUrl: editionUrl,
		})
	})
	return books, nil
}

// GetBooks takes a list and a specific shelf. It returns the books from that shelf and an error.
// Example url: /review/list/68156753
func GetBooks(hostUrl, goodreadsList, shelf string) ([]Book, error) {
	books := []Book{}
	for p := 1; ; p++ {
		url := fmt.Sprintf("%s%s?page=%d&per_page=20&shelf=%s", hostUrl, goodreadsList, p, shelf)
		b, err := getBooksFromPage(url)
		if len(b) == 0 {
			break
		}
		books = append(books, b...)
		if err != nil {
			return books, fmt.Errorf("error getting books from %s: %w", url, err)
		}
		// time.Sleep(time.Second) // wait one second, to limit queries to goodreads (not sure if this is necessary)
	}
	return books, nil
}

// getEditionsFromPage takes an url containing editions on goodreads, scapes and returns the editions and an error.
// Example url: fmt.Sprintf("https://www.goodreads.com/work/editions/%s?page=1&per_page=10", editionsUrl)
func getEditionsFromPage(url string) ([]Edition, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("getting books from %s failed: %v", url, err)
	}
	defer resp.Body.Close()

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return []Edition{}, fmt.Errorf("failed to get editions from '%s' failed: %w", url, err)
	}

	var editions []Edition

	// Regex to clean up the ISBN, removing anything within parentheses
	reISBN := regexp.MustCompile(`\s*\(.*?\)`)

	// Find each "elementList clearFix"
	doc.Find(".elementList.clearFix").Each(func(i int, selection *goquery.Selection) {
		var e Edition

		// Extract format (e.g., "Hardcover, 396 pages")
		selection.Find(".dataRow").Each(func(j int, dataRow *goquery.Selection) {
			text := strings.TrimSpace(dataRow.Text())
			if strings.Contains(text, "pages") {
				// Take the part before the comma
				e.Format = strings.Split(text, ",")[0]
			}
		})

		// Extract ISBN
		selection.Find(".dataTitle:contains('ISBN')").Each(func(j int, dataTitle *goquery.Selection) {
			dataValue := dataTitle.SiblingsFiltered(".dataValue").Text()
			cleanedISBN := strings.TrimSpace(reISBN.ReplaceAllString(dataValue, ""))
			e.Isbn = cleanedISBN
		})

		// Extract language (Edition language)
		selection.Find(".dataTitle:contains('Edition language')").Each(func(j int, dataTitle *goquery.Selection) {
			dataValue := dataTitle.SiblingsFiltered(".dataValue").Text()
			e.Language = strings.TrimSpace(dataValue)
		})

		// Add to editions slice if at least one field is populated
		if e.Format != "" || e.Isbn != "" || e.Language != "" {
			editions = append(editions, e)
		}
	})
	return editions, nil
}

// GetEditions takes an url to editions on goodreads, the required formats and languages. It returns all editions that belong to that work, adhere to the parameters and have a ISBN.
// Example expected url: "/work/editions/94024291"
func GetEditions(hostUrl, editionsUrl string, formats []string, languages []string) ([]Edition, error) {
	editions := []Edition{}
	for p := 1; ; p++ {
		url := fmt.Sprintf("%s%s?page=%d&per_page=10", hostUrl, editionsUrl, p)

		// Get all editions from page
		results, err := getEditionsFromPage(url)
		if len(results) == 0 {
			break
		}

		// Filter to include only editions that have the required format, language and where isbn is not empty
		for _, e := range results {
			if contains(formats, e.Format) && contains(languages, e.Language) && e.Isbn != "" {
				editions = append(editions, e)
			}
		}

		if err != nil {
			return editions, fmt.Errorf("error getting editions from %s: %w", url, err)
		}
		// time.Sleep(time.Second) // wait one second, to limit queries to goodreads (not sure if this is necessary)
	}
	return editions, nil
}

// contains takes a slice of string an value, checks if that value exist in the slice and returns true if it does.
func contains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

// getEditionsUrl takes an url goodreads, the url to the work and retrieves the EditionsUrl from that work.
// E.g. for hostUrl "goodreads.com" and workUrl "/book/show/45047384", the corresponding editionsUrl is "/work/editions/62945242".
func getEditionUrl(hostUrl, workUrl string) (string, error) {
	url := fmt.Sprint(hostUrl, workUrl)
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("getting edition url from %s failed: %w", url, err)
	}
	defer resp.Body.Close()

	// Read the body.
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", fmt.Errorf("creating document failed: %w", err)
	}

	var editionUrl string

	// Define a regex pattern to extract the editionURL
	baseUrl := "/work/editions/"
	pattern := regexp.MustCompile(`/work/quotes/(\d+)`)

	// Iterate through all <a> tags
	doc.Find("a[href]").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists {
			// Check if the href matches the pattern
			matches := pattern.FindStringSubmatch(href)
			if len(matches) > 1 {
				editionUrl = fmt.Sprint(baseUrl, matches[1])
			}
		}
	})

	return editionUrl, nil
}

// NewLibrary takes all parameters for the url (host, list and shelf), the required book formats and languages. It returns an empty library with all parameters.
func NewLibrary(listUrl, shelf string, formats []string, languages []string) Library {
	l := Library{
		hostUrl:   HostUrl,
		ListUrl:   listUrl,
		Shelf:     shelf,
		Formats:   formats,
		Languages: languages,
	}
	return l
}

// Update connects to the host with all parameters that are specified in the type. It retrieves all books and returns an error.
func (l *Library) Update() error {
	var err error
	l.Books, err = GetBooks(l.hostUrl, l.ListUrl, l.Shelf)
	for i := range l.Books {
		// get the Editions Url
		l.Books[i].EditionsUrl, _ = getEditionUrl(l.hostUrl, l.Books[i].WorkUrl)
		// get all editions from that url
		l.Books[i].Editions, err = GetEditions(HostUrl, l.Books[i].EditionsUrl, l.Formats, l.Languages)
	}
	return err
}
