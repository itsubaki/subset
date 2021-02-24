package subset

import "math/rand"

func Select(backendIDs []int, clientID int, subsetSize int) []int {
	subsetCount := len(backendIDs) / subsetSize

	round := clientID / subsetCount
	rand.Seed(int64(round))
	shuffled := Shuffle(backendIDs, round)

	subsetID := clientID % subsetCount
	start := subsetID * subsetSize

	return shuffled[start : start+subsetSize]
}

func Shuffle(backendIDs []int, round int) []int {
	shuffled := make([]int, len(backendIDs))
	copy(shuffled, backendIDs)

	for i := 0; i < len(shuffled); i++ {
		j := rand.Intn(i + 1)
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	}

	return shuffled
}
