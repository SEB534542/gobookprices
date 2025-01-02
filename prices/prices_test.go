package gobookprices

import (
	_ "embed"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	//go:embed testdata/bol1.html
	mockHtmlBol string
	//go:embed testdata/kobo1.html
	mockHtmlKobo string
	//go:embed testdata/ob1.html
	mockHtmlOB1 string
	//go:embed testdata/ob2.html
	mockHtmlOB2 string
)

func TestGetPriceBol(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(mockHtmlBol))
	}))
	defer server.Close()

	isbn := "9789021462691"
	want := 7.99
	got, err := getPriceBol(server.URL, isbn)

	switch {
	case err != nil:
		t.Errorf("error getting price for '%s': %s\nWant: '%.2f'\tGot: '%.2f'\n", isbn, err, want, got)
	case want != got:
		t.Errorf("Want: '%.2f'\tGot: '%.2f'", want, got)
	}
}

func TestGetPriceKobo(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(mockHtmlKobo))
	}))
	defer server.Close()

	isbn := "9789021462691"
	want := 7.99
	// got, err := getPriceKobo(HostKobo, isbn)
	got, err := getPriceKobo(server.URL, isbn)

	switch {
	case err != nil:
		t.Errorf("error getting price for '%s': %s\nWant: '%.2f'\tGot: '%.2f'\n", isbn, err, want, got)
	case want != got:
		t.Errorf("Want: '%.2f'\tGot: '%.2f'", want, got)
	}
}

func TestGetPriceOB(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.RawQuery, "9789021462691") {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(mockHtmlOB1))
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(mockHtmlOB2))
		}
	}))
	defer server.Close()

	t.Run("default test", func(t *testing.T) {
		isbn := "9789021462691"
		want := 0.00
		got, err := getPriceOB(HostOB, isbn)

		switch {
		case err != nil:
			t.Errorf("error getting price for '%s': %s\nWant: '%.2f'\tGot: '%.2f'\n", isbn, err, want, got)
		case want != got:
			t.Errorf("Want: '%.2f'\tGot: '%.2f'", want, got)
		}
	})

	t.Run("test with no results", func(t *testing.T) {
		isbn := "97890214626918989"
		want := 0.00
		got, err := getPriceOB(HostOB, isbn)

		if err != ErrorNotFound || want != got {
			switch {
			case err != ErrorNotFound:
				t.Errorf("error getting price for '%s': %s\nWant: '%.2f'\tGot: '%.2f'\n", isbn, err, want, got)
			case err != nil:
				t.Errorf("error getting price for '%s': %s\nWant: '%.2f'\tGot: '%.2f'\n", isbn, err, want, got)
			case want != got:
				t.Errorf("Want: '%.2f'\tGot: '%.2f'", want, got)
			}
		}
	})
}