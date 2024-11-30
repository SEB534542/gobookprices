package gobookprices

import (
	"reflect"
	"testing"
)

func TestGetWorks(t *testing.T) {
	// TODO: test for editions with multiple works?
	want := "/works/OL16807297W"
	got, err := getWork("9780575104181")
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
		want:= []string{} // no results expected
		got, err := getEditions("/works/OL5623531W", []string{"/languages/eng"})
		switch {
		case err != nil:
			t.Errorf("Unknown error when fetching works for %v: %v", want, err)
		case !reflect.DeepEqual(want, got):
			t.Errorf("\nWant: '%+v', Got: '%+v'", want, got)
		}
	})
}
