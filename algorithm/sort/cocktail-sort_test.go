package sort

import (
	"testing"
)

func TestCocktailSort(t *testing.T) {

	Test(CocktailSort, t)
}

func BenchmarkCocktailSort(b *testing.B) {

	Benchmark(CocktailSort, b)
}
