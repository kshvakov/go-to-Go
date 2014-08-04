package sort

import (
	"testing"
)

func TestSelectionSort(t *testing.T) {

	Test(SelectionSort, t)
}

func BenchmarkSelectionSort(b *testing.B) {

	Benchmark(SelectionSort, b)
}
