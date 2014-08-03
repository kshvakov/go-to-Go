package tree

import (
	"errors"
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
	//@todo: убрать рекурсивный вызов
	if root.Key > node.Key {

		root.Left = b.add(root.Left, node)

	} else {

		root.Right = b.add(root.Right, node)
	}

	return root
}

func (b *BinaryTree) Get(key int) (string, error) {

	root := b.root

	for root != nil {

		if root.Key == key {

			return root.Value, nil
		}

		if root.Key > key {

			root = root.Left

		} else {

			root = root.Right
		}
	}

	return "", errors.New("not found")
}

// @todo: упростить через tmp
func (b *BinaryTree) Delete(key int) {

	var parent *BinaryNode

	current := b.root

	for current != nil {

		if current.Key > key {

			parent = current

			current = current.Left

		} else if current.Key < key {

			parent = current

			current = current.Right

		} else {

			if current.Left == nil && current.Right == nil {

				if parent != nil {

					if parent.Left == current {

						parent.Left = nil

					} else {

						parent.Right = nil
					}
				} else {

					b.root = nil
				}

				return
			}

			if current.Left != nil && current.Right != nil {

				if parent != nil {

					if parent.Left == current {

						parent.Left = nil

					} else {

						parent.Right = nil
					}

					b.add(parent, current.Left)
					b.add(parent, current.Right)

				} else if current == b.root {

					b.root = current.Right

					b.add(b.root, current.Left)
				}

				return
			}

			if current.Left != nil {

				if parent != nil {

					if parent.Left == current {

						parent.Left = nil

					} else {

						parent.Right = nil
					}

					b.add(parent, current.Left)

				} else {

					b.root = current.Left
				}

				return
			}

			if current.Right != nil {

				if parent != nil {

					if parent.Left == current {

						parent.Left = nil

					} else {

						parent.Right = nil
					}

					b.add(parent, current.Right)

				} else {

					b.root = current.Right
				}

				return
			}
		}
	}
}

func (b *BinaryTree) Min() (int, error) {

	root := b.root

	for root != nil {

		if root.Left == nil {

			return root.Key, nil
		}

		root = root.Left
	}

	return -1, errors.New("not found")
}

func (b *BinaryTree) Max() (int, error) {

	root := b.root

	for root != nil {

		if root.Right == nil {

			return root.Key, nil
		}

		root = root.Right
	}

	return -1, errors.New("not found")
}
