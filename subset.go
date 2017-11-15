package subset

import (
	"math/rand"
)

type Backend struct {
	ID int
}

type Subset struct {
	Shuffled   map[int][]*Backend
	backends   []*Backend
	subsetSize int
	count      int
}

func NewSubset(backends []*Backend, subsetSize int) *Subset {
	return &Subset{
		Shuffled:   make(map[int][]*Backend),
		backends:   backends,
		subsetSize: subsetSize,
		count:      len(backends) / subsetSize,
	}
}

func (s *Subset) Select(clientID int) []*Backend {
	round := clientID / s.count
	shuffled := s.Shuffle(round)
	subsetID := clientID % s.count
	start := subsetID * s.subsetSize
	return shuffled[start : start+s.subsetSize]
}

func (s *Subset) Shuffle(round int) []*Backend {
	if backends, ok := s.Shuffled[round]; ok {
		return backends
	}

	rand.Seed(int64(round))
	shuffled := s.backends
	for i := len(shuffled) - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	}

	s.Shuffled[round] = shuffled
	return shuffled
}
