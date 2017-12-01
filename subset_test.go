package subset

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSubset(t *testing.T) {
	backendSize := 300
	clientSize := 300
	subsetSize := 30

	backends := []Backend{}
	for i := 0; i < backendSize; i++ {
		backends = append(backends, Backend{ID: i})
	}

	clients := []Client{}
	for i := 0; i < clientSize; i++ {
		clients = append(clients, Client{ID: i})
	}

	selected := [][]Backend{}
	for _, c := range clients {
		s := Subset(backends, c.ID, subsetSize)
		selected = append(selected, s)
	}

	// check dup
	for i := 0; i < len(selected); i++ {
		for j := 0; j < len(selected); j++ {
			if i == j {
				continue
			}

			if reflect.DeepEqual(selected[i], selected[j]) {
				fmt.Println(selected[i])
				fmt.Println(selected[j])
				t.Error(selected)
			}
		}
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

	for i, count := range stats {
		if count == 30 {
			continue
		}
		t.Error(stats[i])
	}
}
