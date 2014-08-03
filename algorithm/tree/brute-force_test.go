package tree

import (
	"errors"
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

	return "", errors.New("not found")
}

func (b *BruteForce) Min() (int, error) {

	if len(b.list) != 0 {

		min := b.list[0].Key

		for _, N := range b.list {

			if N.Key < min {

				min = N.Key
			}
		}

		return min, nil
	}

	return 0, errors.New("not found")
}

func (b *BruteForce) Max() (int, error) {

	if len(b.list) != 0 {

		max := b.list[0].Key

		for _, N := range b.list {

			if N.Key > max {

				max = N.Key
			}
		}

		return max, nil
	}

	return 0, errors.New("not found")
}

func (b *BruteForce) Delete(key int) {

	for i, N := range b.list {

		if N.Key == key {

			b.list[i], b.list = b.list[len(b.list)-1], b.list[:len(b.list)-1]

			return
		}
	}
}

type BruteForceNode struct {
	Key   int
	Value string
}

func TestBruteForceMin(t *testing.T) {

	tree := new(BruteForce)

	for _, i := range rand.Perm(528) {

		tree.Set(i, fmt.Sprintf("format %d", i))
	}

	v, err := tree.Min()

	if err != nil {

		t.Error(err.Error())
	}

	if v != 0 {
		t.Error("v != 0")
	}
}

func TestBruteForceMax(t *testing.T) {

	tree := new(BruteForce)

	for _, i := range rand.Perm(528) {

		tree.Set(i, fmt.Sprintf("format %d", i))
	}

	v, err := tree.Max()

	if err != nil {

		t.Error(err.Error())
	}

	if v != 527 {
		t.Error("v != 0")
	}
}

func TestBruteForceDelete(t *testing.T) {
	tree := new(BruteForce)

	for _, i := range rand.Perm(528) {

		tree.Set(i, fmt.Sprintf("format_%d", i))
	}

	v, err := tree.Get(23)

	if err != nil {

		t.Error(err.Error())
	}

	if v != "format_23" {

		t.Error("v != format_23")
	}

	tree.Delete(23)

	if _, err := tree.Get(23); err == nil {
		t.Error("key found")
	}
}

func BenchmarkBruteForceSet(b *testing.B) {

	tree := new(BruteForce)

	for _, i := range rand.Perm(b.N) {

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

	for _, i := range rand.Perm(items) {

		tree.Set(i, fmt.Sprintf("format %d", i))
	}

	b.StartTimer()

	for t := 0; t <= b.N; t++ {

		tree.Get(t)
	}
}
