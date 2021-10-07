// Package entry reads and writes to []Entry
package passentry

import (
	"fmt"
	"strings"
)

type Entries []Entry

type Entry struct {
	Url      string
	Username string
	Password string
	Extra    string
	Name     string
	Grouping string
	Fav      bool
}

func (e *Entries) SearchByName(value string) Entries {
	entries := Entries{}
	for _, entry := range *e {
		if strings.Contains(strings.ToLower(entry.Name), strings.ToLower(value)) {
			entries = append(entries, entry)
		}
	}
	return entries
}

func (e *Entries) SearchByGroup(value string) Entries {
	entries := Entries{}
	for _, entry := range *e {
		if strings.ToLower(entry.Grouping) == strings.ToLower(value) {
			entries = append(entries, entry)
		}
	}
	return entries
}

func (e *Entries) Groups() map[string]int {
	groups := make(map[string]int)
	for _, entry := range *e {
		groups[entry.Grouping]++
	}
	return groups
}

func (e *Entries) Insert(entry Entry, list *Entries) {
	*list = append(*list, entry)
}

func DisplayEntry(entry Entry) {
	fmt.Printf("Name: %s\nUsername: %s\nPassword: %s\n",
		entry.Name, entry.Username, entry.Password)
}
