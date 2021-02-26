package subset

import "math/rand"

type Subset struct {
	ids  []int
	size int
}

func New(ids []int, size int) *Subset {
	return &Subset{
		ids:  ids,
		size: size,
	}
}

func (s *Subset) Select(id int) []int {
	subsetCount := len(s.ids) / s.size

	round := id / subsetCount
	rand.Seed(int64(round))
	shuffled := Shuffle(s.ids, round)

	subsetID := id % subsetCount
	start := subsetID * s.size

	return shuffled[start : start+s.size]
}

func Shuffle(ids []int, round int) []int {
	shuffled := make([]int, len(ids))
	copy(shuffled, ids)

	for i := 0; i < len(shuffled); i++ {
		j := rand.Intn(i + 1)
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	}

	return shuffled
}
