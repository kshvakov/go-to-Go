package search

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {

	Test(BinarySearch, t)
}

func BenchmarkBinarySearch(b *testing.B) {

	Benchmark(BinarySearch, b)
}
