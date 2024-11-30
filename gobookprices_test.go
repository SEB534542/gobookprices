package gobookprices

import (
	"reflect"
	"testing"
)

func TestGetWorks(t *testing.T) {
	// TODO: test for editions with multiple works?
	want := "/works/OL16807297W"
	got, err := findWork("9780575104181")
	switch {
	case err != nil:
		t.Errorf("Unknown error when fetching works for %v: %v", want, err)
	case want != got:
		t.Errorf("Want: '%+v', Got: '%+v'", want, got)
	}
}

func TestGetEditions(t *testing.T) {
	// TODO: test with editions that have multiple ISBN?
	t.Run("Get edition", func(t *testing.T) {
		want := []string{
			"9033405474",
		}
		got, err := getEditions("/works/OL5623531W", []string{"/languages/eng", "/languages/dut"})
		switch {
		case err != nil:
			t.Errorf("Unknown error when fetching works for %v: %v", want, err)
		case !reflect.DeepEqual(want, got):
			t.Errorf("\nWant: '%+v', Got: '%+v'", want, got)
		}
	})

	t.Run("Zero results test", func(t *testing.T) {
		want := []string{} // no results expected
		got, err := getEditions("/works/OL5623531W", []string{"/languages/eng"})
		switch {
		case err != nil:
			t.Errorf("Unknown error when fetching works for %v: %v", want, err)
		case !reflect.DeepEqual(want, got):
			t.Errorf("\nWant: '%+v', Got: '%+v'", want, got)
		}
	})
}

func TestGetWork(t *testing.T) {
	want := Work{
		Title:     "Menselijke levensloop",
		Author:    "Hendrik Cammaer",
		Url:       "/works/OL5623531W",
		Languages: []string{"/languages/eng", "/languages/dut"},
	}
	got, err := GetWork("9033405474", []string{"/languages/eng", "/languages/dut"})
	switch {
	case err != nil:
		t.Errorf("Unknown error when fetching works for %v: %v", want, err)
	case !reflect.DeepEqual(want, got):
		t.Errorf("\nWant: '%+v', Got: '%+v'", want, got)
	}

	// Get Editions
	got.GetEditions()
	want.Editions = []Edition{
		{
			Title:    "Menselijke levensloop",
			Isbn:     "9033405474",
			Language: "/languages/dut",
		}}
	switch {
	case err != nil:
		t.Errorf("Unknown error when fetching works for %v: %v", want, err)
	case !reflect.DeepEqual(want, got):
		t.Errorf("\nWant: '%+v', Got: '%+v'", want, got)
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
