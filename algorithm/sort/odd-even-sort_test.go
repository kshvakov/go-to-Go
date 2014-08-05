package sort

import (
	"testing"
)

func TestOddEvenSort(t *testing.T) {

	Test(OddEvenSort, t)
}

func BenchmarkOddEvenSort(b *testing.B) {

	Benchmark(OddEvenSort, b)
}
