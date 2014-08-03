package tree

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestBinaryTreeSetGet(t *testing.T) {

	tree := new(BinaryTree)

	tree.Set(1, "a")
	tree.Set(2, "b")
	tree.Set(3, "c")

	value, err := tree.Get(2)

	if err != nil {

		t.Error(err.Error())
	}

	if value != "b" {

		t.Error(fmt.Sprintf("%s != b", value))
	}

	tree.Set(2, "b_2")

	value, err = tree.Get(2)

	if err != nil {

		t.Error(err.Error())
	}

	if value != "b_2" {

		t.Error(fmt.Sprintf("%s != b", value))
	}

	for _, i := range rand.Perm(2000) {

		if i > 5 {

			tree.Set(i, fmt.Sprintf("value_%d", i))
		}
	}

	value, err = tree.Get(122)

	if err != nil {

		t.Error(err.Error())
	}

	if value != "value_122" {

		t.Error(fmt.Sprintf("%s != value_122", value))
	}
}

func BenchmarkBinaryTreeGet5000(b *testing.B) {
	BinaryTreeGet(b, 5000)
}

func BenchmarkBinaryTreeGet10000(b *testing.B) {
	BinaryTreeGet(b, 10000)
}

func BenchmarkBinaryTreeGet20000(b *testing.B) {
	BinaryTreeGet(b, 20000)
}

func BenchmarkBinaryTreeGet40000(b *testing.B) {
	BinaryTreeGet(b, 40000)
}

func BinaryTreeGet(b *testing.B, items int) {

	tree := new(BinaryTree)

	b.StopTimer()

	for _, i := range rand.Perm(items) {

		tree.Set(i, fmt.Sprintf("value %d", i))
	}

	b.StartTimer()

	for t := 0; t <= b.N; t++ {

		tree.Get(t)
	}
}

func BenchmarkBinaryTreeSet(b *testing.B) {

	tree := new(BinaryTree)

	for _, i := range rand.Perm(b.N) {

		tree.Set(i, fmt.Sprintf("value %d", i))
	}
}
