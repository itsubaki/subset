package subset

import (
	"math/rand"
)

func Subset(backends []Backend, clientID int, subsetSize int) []Backend {
	subsetCount := len(backends) / subsetSize

	round := clientID / subsetCount
	rand.Seed(int64(round))
	shuffled := Shuffle(backends, round)

	subsetID := clientID % subsetCount
	start := subsetID * subsetSize

	return shuffled[start : start+subsetSize]
}

func Shuffle(backends []Backend, round int) []Backend {
	shuffled := make([]Backend, len(backends))
	copy(shuffled, backends)

	for i := 0; i < len(shuffled); i++ {
		j := rand.Intn(i + 1)
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	}

	return shuffled
}
