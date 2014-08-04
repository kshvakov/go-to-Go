package search

import (
	"math/rand"
	"testing"
)

func TestBinarySearch(t *testing.T) {

	haystack := []int{2, 5, 7, 9, 20, 21, 42}

	if val := BinarySearch(9, haystack); val == -1 {

		t.Error("not found")
	}

	if val := BinarySearch(42, haystack); val == -1 {

		t.Error("not found")
	}

	if val := BinarySearch(2, haystack); val == -1 {

		t.Error("not found")
	}

	if val := BinarySearch(19, haystack); val != -1 {

		t.Error("found")
	}
}

func BenchmarkBinarySearch(b *testing.B) {

	var haystack []int

	b.StopTimer()

	for i := 0; i < b.N; i++ {

		haystack = append(haystack, i)
	}

	needles := rand.Perm(b.N)

	b.StartTimer()

	for n := range needles {

		if v := BinarySearch(n, haystack); v == -1 {

			b.Fatalf("%d not found", n)
		}
	}
}
