package ravg

import (
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

// RAvg keeps a running or rolling average that is computed over a number of
// samples of a type of Number.
type RAvg[T Number] struct {
	samples []T
	next    int
	full    bool
}

// New creaates a new RAvg instance that stores the specified number of samples
// of a type of Number.
func New[T Number](size int) *RAvg[T] {
	return &RAvg[T]{
		samples: make([]T, size),
	}
}

// Len returns the number of samples the average is computed over.
func (r *RAvg[T]) Len() int {
	return len(r.samples)
}

// Put adds a new sample to the sample set, replacing the oldest.
func (r *RAvg[T]) Put(sample T) {
	r.samples[r.next] = sample
	r.next++
	if r.next == len(r.samples) {
		r.next = 0
		r.full = true
	}
}

// Mean returns the mean value of the samples.
func (r *RAvg[T]) Mean() T {
	size, sum := r.mean()
	return sum / T(size)
}

// FMean returns the mean value of the samples as a float64.
func (r *RAvg[T]) FMean() float64 {
	size, sum := r.mean()
	return float64(sum) / float64(size)
}

func (r *RAvg[T]) mean() (int, T) {
	size := len(r.samples)
	if !r.full {
		size = r.next
	}
	var sum T
	for i := 0; i < size; i++ {
		sum += r.samples[i]
	}
	return size, sum
}
