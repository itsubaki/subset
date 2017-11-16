package subset

import (
	"fmt"
	"testing"
)

func TestSubset(t *testing.T) {
	backendSize := 300
	clientSize := 300
	subsetSize := 30

	backends := []Backend{}
	for i := 0; i < backendSize; i++ {
		backends = append(backends, Backend{i})
	}

	selected := [][]Backend{}
	for i := 0; i < clientSize; i++ {
		s := Subset(backends, i, subsetSize)
		selected = append(selected, s)

		fmt.Println(s)
	}

	// count by backendID
	flat := []Backend{}
	for _, b := range selected {
		flat = append(flat, b...)
	}

	stats := make(map[int]int)
	for _, backend := range flat {
		stats[backend.ID] = stats[backend.ID] + 1
	}

	for _, count := range stats {
		if count == 30 {
			continue
		}
		t.Error(stats)
	}
}
