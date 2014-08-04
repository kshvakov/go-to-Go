package sort

import (
	"testing"
)

func TestCombSort(t *testing.T) {

	Test(CombSort, t)
}

func BenchmarkCombSort(b *testing.B) {

	Benchmark(CombSort, b)
}
