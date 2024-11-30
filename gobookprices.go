package gobookprices

import (
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
)

type Work struct {
	Title    string    // Title of the work
	Author   string    // Author(s)
	Editions []Edition // Contains all editions identified
	Url      string    // Url to the work at openlibrary
}

type Edition struct {
	Title    string // Title of the edition
	Isbn     string // Isbn13, or if not found the Isbn10
	Language string // Language of the edition
	Ebook    bool   // If the edition is a Ebook
}

// findWork takes an ISBN, looks up the edition on openlibrary and returns the first work.
func findWork(isbn string) (string, error) {
	url := fmt.Sprintf("https://openlibrary.org/isbn/%s.json", isbn)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result struct {
		Works []struct {
			Key string `json:"key"`
		} `json:"works"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}
	return result.Works[0].Key, nil
}

// getEditions takes a reference to a work on openlibrary and the required languages.
// It returns all IBNS13' that belong to that work. If no ISBN13 exists, it will return the ISBN10 for that edition.
func getEditions(work string, langs []string) ([]string, error) {
	url := fmt.Sprintf("https://openlibrary.org/%s/editions.json", work)
	resp, err := http.Get(url)
	if err != nil {
		return []string{}, err
	}
	defer resp.Body.Close()

	var result struct {
		Entries []struct {
			FullTitle string   `json:"full_title,omitempty"`
			Isbn13    []string `json:"isbn_13"`
			Isbn10    []string `json:"isbn_10"`
			Languages []struct {
				Key string `json:"key"`
			} `json:"languages,omitempty"`
			PhysicalFormat string   `json:"physical_format,omitempty"`
			Series         []string `json:"series,omitempty"`
		}
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return []string{}, err
	}

	// TODO: filter on languages and return only that list
	editions := make([]string, 0, len(result.Entries))
	for _, v := range result.Entries {
		// check if edition is in requested language
		for _, lang := range v.Languages {
			if slices.Contains(langs, lang.Key) {
				var isbns []string
				switch {
				case len(v.Isbn13) != 0:
					isbns = append(isbns, v.Isbn13...)
				case len(v.Isbn10) != 0:
					isbns = append(isbns, v.Isbn10...)
				}
				editions = append(editions, isbns...)
				break
			}
		}
	}
	return editions, nil
}

func getWork(isbn string) (Work, error) {
	work := Work{}
	var err error

	// Get Work url
	work.Url, err = findWork(isbn)
	if err != nil {
		return work, err // Todo: rewrite this to proper error usage
	}
	// Get Work details
	url := fmt.Sprintf("https://openlibrary.org/%s.json", work.Url)
	resp, err := http.Get(url)
	if err != nil {
		return work, err
	}
	defer resp.Body.Close()

	var result struct {
		Title   string `json:"title"`
		Authors []struct {
			Author struct {
				Key string `json:"key"`
			} `json:"author"`
		} `json:"authors"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return work, err
	}

	work.Title = result.Title
	// get Author details 
	l := len(result.Authors)
	switch {
	case l == 1:
		work.Author, err = getAuthorDetails(result.Authors[0].Author.Key)
	case l > 1:
		work.Author,err = getAuthorDetails(result.Authors[0].Author.Key)
		for _, a := range result.Authors[1:] {
			var author string
			author, err = getAuthorDetails(fmt.Sprint(a)) // TODO: err gets overwritten, add?
			work.Author += fmt.Sprintf(", %s", author)
		}
	}
	return work, err
}

// getAuthorDetails takes the openlibrary reference to an author and returns the name of the author as a string.
func getAuthorDetails(author string) (string, error) {
	url := fmt.Sprintf("https://openlibrary.org/%s.json", author)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result struct {
		Name string `json:"name"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	return result.Name, nil
}
