package tree

import (
	"errors"
	"fmt"
	"math/rand"
	"testing"
)

type Linear struct {
	list []*LinearNode
}

func (b *Linear) Set(key int, value string) {
	b.list = append(b.list, &LinearNode{Key: key, Value: value})
}

func (b *Linear) Get(key int) (string, error) {

	for _, N := range b.list {

		if N.Key == key {

			return N.Value, nil
		}
	}

	return "", errors.New("not found")
}

func (b *Linear) Min() (int, error) {

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

func (b *Linear) Max() (int, error) {

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

func (b *Linear) Delete(key int) {

	for i, N := range b.list {

		if N.Key == key {

			b.list[i], b.list = b.list[len(b.list)-1], b.list[:len(b.list)-1]

			return
		}
	}
}

type LinearNode struct {
	Key   int
	Value string
}

func TestLinearMin(t *testing.T) {

	tree := new(Linear)

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

func TestLinearMax(t *testing.T) {

	tree := new(Linear)

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

func TestLinearDelete(t *testing.T) {
	tree := new(Linear)

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

func BenchmarkLinearSet(b *testing.B) {

	tree := new(Linear)

	for _, i := range rand.Perm(b.N) {

		tree.Set(i, fmt.Sprintf("format %d", i))
	}
}

func BenchmarkLinearGet5000(b *testing.B) {
	LinearGet(b, 5000)
}

func BenchmarkLinearGet10000(b *testing.B) {
	LinearGet(b, 10000)
}

func BenchmarkLinearGet20000(b *testing.B) {
	LinearGet(b, 20000)
}

func BenchmarkLinearGet40000(b *testing.B) {
	LinearGet(b, 40000)
}

func LinearGet(b *testing.B, items int) {

	b.StopTimer()

	tree := new(Linear)

	for _, i := range rand.Perm(items) {

		tree.Set(i, fmt.Sprintf("format %d", i))
	}

	b.StartTimer()

	for t := 0; t <= b.N; t++ {

		tree.Get(t)
	}
}
