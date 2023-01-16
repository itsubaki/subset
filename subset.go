package subset

import "math/rand"

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
	out := make([]int, len(s.ids))
	copy(out, s.ids)

	rand.Seed(int64(round))
	for i := 0; i < len(out); i++ {
		j := rand.Intn(i + 1)
		out[i], out[j] = out[j], out[i]
	}

	return out
}
