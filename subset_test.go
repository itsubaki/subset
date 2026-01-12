package subset_test

import (
	"fmt"

	"github.com/itsubaki/subset"
)

func testdatagen(size int) []int {
	ids := make([]int, size)
	for i := range size {
		ids[i] = i
	}

	return ids
}

func Example() {
	backendSize := 100
	backendIDs := testdatagen(backendSize)

	subsetSize := 10
	s := subset.New(backendIDs, subsetSize)

	clientSize := 10
	clientIDs := testdatagen(clientSize)
	for _, clientID := range clientIDs {
		selected := s.Select(clientID)
		fmt.Printf("ClientID: %v -> BackendIDs: %2v\n", clientID, selected)
	}

	// Output:
	// ClientID: 0 -> BackendIDs: [55 58 54 75 66 33 34 24  9 16]
	// ClientID: 1 -> BackendIDs: [56 14 31 63 88 60 46 47 94 39]
	// ClientID: 2 -> BackendIDs: [45  1 12 26 21 52 32  7 74 70]
	// ClientID: 3 -> BackendIDs: [93 64 96 37 99 90 11  6 27 87]
	// ClientID: 4 -> BackendIDs: [83 17 76  0 35 95 43 92 77 91]
	// ClientID: 5 -> BackendIDs: [36 97 73 28 19 30 48 10 61 82]
	// ClientID: 6 -> BackendIDs: [ 3 59 40 25 44 13 86  4 98 72]
	// ClientID: 7 -> BackendIDs: [20 18 41  5 49 84 68 81 51 38]
	// ClientID: 8 -> BackendIDs: [29  2 67 15 71 89 57 23 78 79]
	// ClientID: 9 -> BackendIDs: [62 50 85 65 42  8 53 69 22 80]
}
