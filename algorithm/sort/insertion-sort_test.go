package sort

import (
	"testing"
)

func TestInsertionSort(t *testing.T) {

	Test(InsertionSort, t)
}

func BenchmarkInsertionSort(b *testing.B) {

	Benchmark(InsertionSort, b)
}
