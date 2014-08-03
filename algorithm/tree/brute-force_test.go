package tree

import (
	"fmt"
	"math/rand"
	"testing"
)

type BruteForce struct {
	list []*BruteForceNode
}

func (b *BruteForce) Set(key int, value string) {
	b.list = append(b.list, &BruteForceNode{Key: key, Value: value})
}

func (b *BruteForce) Get(key int) (string, error) {

	for _, N := range b.list {

		if N.Key == key {

			return N.Value, nil
		}
	}

	return "", fmt.Errorf("%d not found", key)
}

type BruteForceNode struct {
	Key   int
	Value string
}

func BenchmarkBruteForceSet(b *testing.B) {

	tree := new(BruteForce)

	for i := range rand.Perm(b.N) {

		tree.Set(i, fmt.Sprintf("format %d", i))
	}
}

func BenchmarkBruteForceGet5000(b *testing.B) {
	BruteForceGet(b, 5000)
}

func BenchmarkBruteForceGet10000(b *testing.B) {
	BruteForceGet(b, 10000)
}

func BenchmarkBruteForceGet20000(b *testing.B) {
	BruteForceGet(b, 20000)
}

func BenchmarkBruteForceGet40000(b *testing.B) {
	BruteForceGet(b, 40000)
}

func BruteForceGet(b *testing.B, items int) {

	b.StopTimer()

	tree := new(BruteForce)

	for i := range rand.Perm(items) {

		tree.Set(i, fmt.Sprintf("format %d", i))
	}

	b.StartTimer()

	for t := 0; t <= b.N; t++ {

		tree.Get(t)
	}
}
