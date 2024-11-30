package gobookprices

import (
	"reflect"
	"testing"
)

func TestFindWork(t *testing.T) {
	want := "/works/OL16807297W"
	got, err := findWork("9780575104181")
	switch {
	case err != nil:
		t.Errorf("Unknown error when fetching works for %v: %v", want, err)
	case want != got:
		t.Errorf("Want: '%+v', Got: '%+v'", want, got)
	}
}

func TestGetAuthorDetails(t *testing.T) {
	want := "Hendrik Cammaer"
	got, err := getAuthorDetails("/authors/OL1357372A")
	switch {
	case err != nil:
		t.Errorf("Unknown error when fetching works for %v: %v", want, err)
	case want != got:
		t.Errorf("\nWant: '%+v', Got: '%+v'", want, got)
	}
}

func TestNewWork(t *testing.T) {
	want := Work{
		Title:     "Menselijke levensloop",
		Author:    "Hendrik Cammaer",
		Url:       "/works/OL5623531W",
		Languages: []string{"/languages/eng", "/languages/dut"},
		Editions: []Edition{
			{
				Title:    "Menselijke levensloop",
				Isbn:     "9033405474",
				Language: "/languages/dut",
			}},
	}
	got, err := NewWork("9033405474", []string{"/languages/eng", "/languages/dut"})
	switch {
	case err != nil:
		t.Errorf("Unknown error when fetching works for %v: %v", want, err)
	case !reflect.DeepEqual(want, got):
		t.Errorf("\nWant: '%+v', Got: '%+v'", want, got)
	}
}
