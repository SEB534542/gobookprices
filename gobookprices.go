package gobookprices

import (
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
)

// type Edition struct {
// 	isbn string // ISBN13 for the book, or if no ISBN13 exists, then it is ISBN10
// }

// getWorks takes an ISBN, looks up the edition on openlibrary and returns the first work.
func getWork(isbn string) (string, error) {
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
			Key       string   `json:"key"`
			Languages []struct {
				Key string `json:"key"`
			} `json:"languages,omitempty"`
			PhysicalFormat string   `json:"physical_format,omitempty"`
			ByStatement    string   `json:"by_statement,omitempty"`
			Contributions  []string `json:"contributions,omitempty"`
			Description    struct {
				Type  string `json:"type"`
				Value string `json:"value"`
			} `json:"description,omitempty"`
			DeweyDecimalClass []string `json:"dewey_decimal_class,omitempty"`
			EditionName       string   `json:"edition_name,omitempty"`
			Ocaid             string   `json:"ocaid,omitempty"`
			OclcNumbers       []string `json:"oclc_numbers,omitempty"`
			PublishCountry    string   `json:"publish_country,omitempty"`
			PublishPlaces     []string `json:"publish_places,omitempty"`
			Series            []string `json:"series,omitempty"`
			WorkTitles        []string `json:"work_titles,omitempty"`
			Weight            string   `json:"weight,omitempty"`
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
