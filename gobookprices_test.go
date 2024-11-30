package gobookprices

import "testing"

func TestGetWorks(t *testing.T) {
	want := "/works/OL16807297W"
	got, _ := getWorks("9780575104181")
	if want != got {
		t.Errorf("\nWant: '%+v', Got: '%+v'", want, got)
	}
}
