package tree

import (
	"fmt"
)

type BinaryNode struct {
	Key   int
	Value string

	Left, Right *BinaryNode
}

type BinaryTree struct {
	root *BinaryNode
}

func (b *BinaryTree) Set(key int, value string) {

	node := &BinaryNode{
		Key:   key,
		Value: value,
	}

	if b.root == nil {

		b.root = node

	} else {

		b.add(b.root, node)
	}
}

func (b *BinaryTree) add(root, node *BinaryNode) *BinaryNode {

	if root == nil {

		root = node

		return root
	}

	if root.Key == node.Key {

		root.Value = node.Value

		return root
	}

	if root.Key > node.Key {

		root.Left = b.add(root.Left, node)

	} else {

		root.Right = b.add(root.Right, node)
	}

	return root
}

func (b *BinaryTree) Get(key int) (string, error) {

	root := b.root

	for {

		if root == nil {

			return "", fmt.Errorf("%d not found", key)
		}

		if root.Key == key {

			return root.Value, nil
		}

		if root.Key > key {

			root = root.Left

		} else {

			root = root.Right
		}

	}

	return "", fmt.Errorf("%d not found", key)
}
