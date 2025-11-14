package subset

import (
	"math/rand/v2"
	"slices"
)

type Subset struct {
	ids   []int
	size  int
	count int
}

func New(ids []int, size int) *Subset {
	return &Subset{
		ids:   ids,
		size:  size,
		count: len(ids) / size,
	}
}

func (s *Subset) Select(id int) []int {
	round := id / s.count
	shuffled := s.Shuffle(round)

	start := (id % s.count) * s.size
	return shuffled[start : start+s.size]
}

func (s *Subset) Shuffle(round int) []int {
	rnd := rand.New(rand.NewPCG(uint64(round), 0))

	out := slices.Clone(s.ids)
	for i := range len(out) {
		j := rnd.IntN(i + 1)
		out[i], out[j] = out[j], out[i]
	}

	return out
}
