package sort

import (
	"math/rand"
	"testing"
)

func Test(f func(slice []int), t *testing.T) {

	slice := []int{2, 1, 56, 21, 6, 89, 11, 2, 3, 52, 42, 11}

	f(slice)

	prev := slice[0]

	for i := range slice {

		if i > 0 {

			prev = slice[i-1]
		}

		if prev > slice[i] {

			t.Fatal("slice not sorted")
		}
	}
}

func Benchmark(f func(slice []int), b *testing.B) {

	b.N = 1000

	for i := 0; i < b.N; i++ {

		b.StopTimer()

		slice := rand.Perm(b.N)

		b.StartTimer()

		f(slice)
	}
}
