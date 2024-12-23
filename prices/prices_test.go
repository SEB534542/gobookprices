package gobookprices

import (
	_ "embed"
	"net/http"
	"net/http/httptest"
	"testing"
)

//go:embed testdata/bol1.html
var mockHtmlBol string

func TestGetPrice(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(mockHtmlBol))
	}))

	isbn := "9789021462691"
	want := 7.99
	got, err := getPriceBol(server.URL, isbn)

	switch {
	case err != nil:
		t.Errorf("error getting price for '%s':Want: '%.2f'\tGot: '%.2f'\n", isbn, want, got)
	case want != got:
		t.Errorf("Want: '%.2f'\tGot: '%.2f'", want, got)
	}
}
