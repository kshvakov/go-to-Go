package search

import (
	"testing"
)

func TestInterpolationSearch(t *testing.T) {

	Test(InterpolationSearch, t)
}

func BenchmarkInterpolationSearch(b *testing.B) {

	Benchmark(InterpolationSearch, b)
}
