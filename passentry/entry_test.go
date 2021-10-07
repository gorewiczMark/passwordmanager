package passentry

import (
	"fmt"
	"testing"
)

var e Entries

func TestSeachByName(t *testing.T) {
	e = []Entry{Entry{Name: "Bob", Url: "www.bob.com"}}
	want := fmt.Sprint([]Entry{Entry{Name: "Bob", Url: "www.bob.com"}})
	got := fmt.Sprint(e.SearchByName("Bob"))
	if got != want {
		t.Errorf("Wanted: %s but Got: %s", want, got)

	}
}
