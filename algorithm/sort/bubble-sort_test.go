package sort

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {

	Test(BubbleSort, t)
}

func BenchmarkBuubleSort(b *testing.B) {

	Benchmark(BubbleSort, b)
}
