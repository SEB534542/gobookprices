package gobookprices

import (
	_ "embed"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

var (
	//go:embed testdata/bol1.html
	mockHtmlBol1 string
	//go:embed testdata/bol2.html
	mockHtmlBol2 string
	//go:embed testdata/kobo1.html
	mockHtmlKobo1 string
	//go:embed testdata/kobo2.html
	mockHtmlKobo2 string
	//go:embed testdata/ob1.html
	mockHtmlOB1 string
	//go:embed testdata/ob2.html
	mockHtmlOB2 string
)

func TestGetPrices(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isbn := "9789021462691"
		if strings.Contains(r.URL.Path, HostBol) {
			switch {
			case strings.Contains(r.URL.RawQuery, isbn):
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(mockHtmlBol1))
			default:
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(mockHtmlBol2))
			}
		} else if strings.Contains(r.URL.Path, HostKobo) {
			switch {
			case strings.Contains(r.URL.RawQuery, isbn):
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(mockHtmlKobo1))
			default:
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(mockHtmlKobo2))
			}
		} else if strings.Contains(r.URL.Path, HostOB) {
			switch {
			case strings.Contains(r.URL.RawQuery, isbn):
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(mockHtmlOB1))
			default:
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(mockHtmlOB2))
			}
		}
	}))
	defer server.Close()
	AllSites[0].Host = fmt.Sprintf("%s/%s/", server.URL, HostBol)
	AllSites[1].Host = fmt.Sprintf("%s/%s/", server.URL, HostKobo)
	AllSites[2].Host = fmt.Sprintf("%s/%s/", server.URL, HostOB)

	t.Run("default test", func(t *testing.T) {
		isbn := "9789021462691"
		want := []float64{7.99, 7.99, 0.00}
		got, err := AllSites.GetPrices(isbn)

		if err != nil {
			t.Error("error getting price:", err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Want: '%v', Got: '%v'\n", want, got)
		}
	})

	// case 2
	t.Run("not existing isbn", func(t *testing.T) {
		isbn := "97890234341462691"
		got, err := AllSites.GetPrices(isbn)
		if !errors.Is(err, ErrorNotFound) {
			t.Errorf("retrieved '%v' with error: %s", got, err)
		}
	})
}