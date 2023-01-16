# subset

 - load balancing algorithm
 - see "A Subset Selection Algorithm: Deterministic Subsetting"@[Site Reliability Engineering](https://landing.google.com/sre/book.html)

## Example

```go
func Example() {
	backendSize := 100
	clientSize := 10
	subsetSize := 10

	testdatagen := func(size int) []int {
		ids := make([]int, 0)
		for i := 0; i < size; i++ {
			ids = append(ids, i)
		}

		return ids
	}

	clientIDs := testdatagen(clientSize)
	backendIDs := testdatagen(backendSize)

	s := subset.New(backendIDs, subsetSize)
	for _, clientID := range clientIDs {
		selected := s.Select(clientID)
		fmt.Printf("ClientID: %v -> BackendIDs: %2v\n", clientID, selected)
	}

	// Output:
	// ClientID: 0 -> BackendIDs: [40 35 50 66 44 88  1 52 67 56]
	// ClientID: 1 -> BackendIDs: [21 72 23 34 86 11 42 20 17 64]
	// ClientID: 2 -> BackendIDs: [27 58 43 46 47 45 87 49 74 30]
	// ClientID: 3 -> BackendIDs: [71 83 25 75 39 78 37 70 33 10]
	// ClientID: 4 -> BackendIDs: [91 99  6 79 59 18 53 76 98  3]
	// ClientID: 5 -> BackendIDs: [57 69 84 14  4 16 54 38 81 36]
	// ClientID: 6 -> BackendIDs: [89 29 32 80 48 60 95 13 92 24]
	// ClientID: 7 -> BackendIDs: [31 73 65 90 51 62 77 85 28 61]
	// ClientID: 8 -> BackendIDs: [ 9 63 93 26 55  2 68  5  7 12]
	// ClientID: 9 -> BackendIDs: [41 96 82 94 19  0 22 15 97  8]
}
```
