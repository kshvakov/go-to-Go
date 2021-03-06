package search

import (
	"math/rand"
	"testing"
)

func LinearSearch(needle int, haystack []int) int {

	for k, v := range haystack {

		if v == needle {

			return k
		}
	}

	return -1
}

func BenchmarkLinearSearch(b *testing.B) {

	b.N = 10000

	var haystack []int

	b.StopTimer()

	for i := 0; i < b.N; i++ {

		haystack = append(haystack, i)
	}

	needles := rand.Perm(b.N)

	b.StartTimer()

	for n := range needles {

		if v := LinearSearch(n, haystack); v == -1 {

			b.Fatalf("%d not found", n)
		}
	}
}
