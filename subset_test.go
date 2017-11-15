package subset

import (
	"fmt"
	"sort"
	"testing"
)

func TestSubset(t *testing.T) {
	backendSize := 300
	clientSize := 300
	subsetSize := 30

	backends := []*Backend{}
	for i := 0; i < backendSize; i++ {
		backends = append(backends, &Backend{i})
	}

	subset := NewSubset(backends, subsetSize)

	selected := []*Backend{}
	for i := 0; i < clientSize; i++ {
		s := subset.Select(i)
		for _, b := range s {
			fmt.Print(b)
		}
		fmt.Println()

		selected = append(selected, s...)
	}

	stats := make(map[int]int)
	for _, backend := range selected {
		stats[backend.ID] = stats[backend.ID] + 1
	}

	ids := []int{}
	for k, _ := range stats {
		ids = append(ids, k)
	}
	sort.Ints(ids)

	for _, id := range ids {
		fmt.Printf("%-3d %d\n", id, stats[id])
	}
}
