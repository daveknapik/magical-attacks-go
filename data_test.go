package main

import (
	"testing"
)

func TestUnique(t *testing.T) {
	data := []string{"Finn", "Jake", "BMO", "Marcy", "Peebs", "BMO"}
	itemCounts := make(map[string]int)

	// count items on a deduped version of the data
	for _, v := range unique(data) {
		itemCounts[v]++
	}

	for k, v := range itemCounts {
		if v > 1 {
			t.Errorf("%s has a count of %v but every item in this slice should only appear once", k, v)
		}
	}
}
