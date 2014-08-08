package search

import (
	"math/rand"
	"testing"
)

func Test(f func(needle int, haystack []int) int, t *testing.T) {

	haystack := []int{2, 5, 7, 9, 9, 18, 20, 21, 25, 27, 35, 42}

	if val := f(9, haystack); val == -1 {

		t.Error("not found")
	}

	if val := f(42, haystack); val == -1 {

		t.Error("not found")
	}

	if val := f(2, haystack); val == -1 {

		t.Error("not found")
	}

	if val := f(19, haystack); val != -1 {

		t.Error("found")
	}
}

func Benchmark(f func(needle int, haystack []int) int, b *testing.B) {

	var haystack []int

	b.StopTimer()

	for i := 0; i < b.N; i++ {

		haystack = append(haystack, i)
	}

	needles := rand.Perm(b.N)

	b.StartTimer()

	for n := range needles {

		if v := f(n, haystack); v == -1 {

			b.Fatalf("%d not found", n)
		}
	}
}
