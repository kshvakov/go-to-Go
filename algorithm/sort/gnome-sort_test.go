package sort

import (
	"testing"
)

func TestGnomeSort(t *testing.T) {

	Test(GnomeSort, t)
}

func BenchmarkGnomeSort(b *testing.B) {

	Benchmark(GnomeSort, b)
}
