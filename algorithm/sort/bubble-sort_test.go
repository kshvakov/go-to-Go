package sort

import (
	"math/rand"
	"testing"
)

func TestBubbleSort(t *testing.T) {

	slice := []int{2, 1, 56, 21, 6, 89, 11, 2, 3, 52, 42, 11}

	BubbleSort(slice)

	min := slice[0]

	for i := 0; i < len(slice)-1; i++ {

		if min > slice[i] {

			t.Fatal("slice not sorted")
		}

		min = slice[i]
	}
}

func BenchmarkBuubleSort(b *testing.B) {

	b.N = 1000

	for i := 0; i < b.N; i++ {

		b.StopTimer()

		slice := rand.Perm(b.N)

		b.StartTimer()

		BubbleSort(slice)
	}
}
