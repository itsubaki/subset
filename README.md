# subset

 - load balancing algorithm
 - see "A Subset Selection Algorithm: Deterministic Subsetting"@[Site Reliability Engineering](https://landing.google.com/sre/book.html)

## Example

```go
func Example() {
	backendSize := 100
	clientSize := 10
	subsetSize := 10

	clients := make([]int, 0)
	for i := 0; i < clientSize; i++ {
		clients = append(clients, i)
	}

	backends := make([]int, 0)
	for i := 0; i < backendSize; i++ {
		backends = append(backends, i)
	}

	for _, c := range clients {
		s := subset.Select(backends, c, subsetSize)
		fmt.Printf("%v: %2v\n", c, s)
	}

	// Output:
	// 0: [40 35 50 66 44 88  1 52 67 56]
	// 1: [21 72 23 34 86 11 42 20 17 64]
	// 2: [27 58 43 46 47 45 87 49 74 30]
	// 3: [71 83 25 75 39 78 37 70 33 10]
	// 4: [91 99  6 79 59 18 53 76 98  3]
	// 5: [57 69 84 14  4 16 54 38 81 36]
	// 6: [89 29 32 80 48 60 95 13 92 24]
	// 7: [31 73 65 90 51 62 77 85 28 61]
	// 8: [ 9 63 93 26 55  2 68  5  7 12]
	// 9: [41 96 82 94 19  0 22 15 97  8]
}
```
