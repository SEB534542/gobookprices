package gobookprices

import (
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
)

type Work struct {
	Title     string    // Title of the work
	Author    string    // Author(s)
	Url       string    // Url to the work at openlibrary
	Languages []string  // Languages that editions should have
	Editions  []Edition // Contains all editions identified
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
		return "", fmt.Errorf("unable to get data from url: %w", err)
	}
	defer resp.Body.Close()

	var result struct {
		Works []struct {
			Key string `json:"key"`
		} `json:"works"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("failed to convert results from json: %w", err)
	}
	return result.Works[0].Key, nil
}

// getWork takes an ISBN and preferred languages and returns the associated Work and an error.
func getWork(isbn string, languages []string) (Work, error) {
	work := Work{}
	work.Languages = languages

	var err error

	// Get Work url
	work.Url, err = findWork(isbn)
	if err != nil {
		return work, fmt.Errorf("unable to get data from url: %w", err)
	}
	// Get Work details
	url := fmt.Sprintf("https://openlibrary.org/%s.json", work.Url)
	resp, err := http.Get(url)
	if err != nil {
		return work, fmt.Errorf("unable to get data from url: %w", err)
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
		return work, fmt.Errorf("failed to convert results from json: %w", err)
	}

	work.Title = result.Title
	// get Author details
	l := len(result.Authors)
	switch {
	case l == 1:
		work.Author, err = getAuthorDetails(result.Authors[0].Author.Key)
	case l > 1:
		work.Author, err = getAuthorDetails(result.Authors[0].Author.Key)
		for _, a := range result.Authors[1:] {
			var author string
			author, err = getAuthorDetails(fmt.Sprint(a))
			work.Author += fmt.Sprintf(", %s", author)
		}
	}
	if err != nil {
		err = fmt.Errorf("error retieving author details: %w", err)
	}
	return work, err
}

// getAuthorDetails takes the openlibrary reference to an author and returns the name of the author as a string.
func getAuthorDetails(author string) (string, error) {
	url := fmt.Sprintf("https://openlibrary.org/%s.json", author)
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("unable to get data from url %s: %w", url, err)
	}
	defer resp.Body.Close()

	var result struct {
		Name string `json:"name"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("failed to convert results from json: %w", err)
	}

	return result.Name, nil
}

// NewWork takes an ISBN and preferred languages and returns the associated Work with underlying editions and an error.
func NewWork(isbn string, languages []string) (Work, error) {
	work, err := getWork(isbn, languages)
	if err != nil {
		return work, fmt.Errorf("error while retrieving work: %w", err)
	}
	err = work.GetEditions()
	if err != nil {
		return work, fmt.Errorf("error while retrieving editions: %w", err)
	}
	return work, nil
}

// GetEditions gathers all editions beloning to the work from openlibary in the Work and returns an error.
func (work *Work) GetEditions() error {
	url := fmt.Sprintf("https://openlibrary.org/%s/editions.json", work.Url)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var result struct {
		Entries []struct {
			Title     string   `json:"title"`
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
		return err
	}

	work.Editions = make([]Edition, 0, len(result.Entries))
	for _, v := range result.Entries {
		// check if edition is in requested language
		for _, lang := range v.Languages {
			if slices.Contains(work.Languages, lang.Key) {
				var isbn string
				switch {
				case len(v.Isbn13) != 0:
					isbn = v.Isbn13[0]
				case len(v.Isbn10) != 0:
					isbn = v.Isbn10[0]
				}
				work.Editions = append(work.Editions, Edition{
					Title:    v.Title,
					Isbn:     isbn,
					Language: lang.Key,
				})
				break
			}
		}
	}
	return nil
}
