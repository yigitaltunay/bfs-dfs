package tree

import (
	"errors"
	"fmt"
)

type Tree struct {
	Value  int
	Left   *Tree
	Right  *Tree
	Repeat int
	Parent *Tree
}

func (t *Tree) Insert(v int) {
	switch {
	case v == t.Value:
		t.Repeat++
	case v < t.Value:
		if t.Left == nil {
			t.Left = &Tree{v, nil, nil, 1, t}
			return
		}
		t.Left.Insert(v)
	case v > t.Value:
		if t.Right == nil {
			t.Right = &Tree{v, nil, nil, 1, t}
			return
		}
		t.Right.Insert(v)
	}
}

func (t *Tree) Find(v int) (int, *Tree, bool) {
	switch {
	case v == t.Value:
		return t.Repeat, t, true
	case v < t.Value:
		if t.Left == nil {
			return 0, nil, false
		}
		return t.Left.Find(v)
	case v > t.Value:
		if t.Right == nil {
			return 0, nil, false
		}
		return t.Right.Find(v)
	}
	return 0, nil, false

}

func (t *Tree) Maximum() int {
	if t.Right == nil {
		return t.Value
	}
	return t.Right.Maximum()
}

func (t *Tree) Minimum() int {
	if t.Left == nil {
		return t.Value
	}
	return t.Left.Minimum()
}

func (t Tree) TypeCheck(v interface{}) error {
	switch v.(type) {
	case int:
		return nil
	default:
		return fmt.Errorf("unexpected type %T", v)

	}
}

func (t *Tree) Delete(v int) error {
	_, tree, present := t.Find(v)
	if !present {
		return errors.New("Element not present in the tree")
	}
	return tree.Replace(v)
}

func (t *Tree) Replace(v int) error {
	if t.Right == nil && t.Left == nil {
		t.Parent = nil
		return nil
	}
	switch {
	case t.Right == nil:
		t.Parent.Left = t.Left
	case t.Left == nil:
		t.Parent.Right = t.Right
	default:
		min := t.Right.Minimum()
		t.Value = min
		t.Right.Delete(min)
	}
	return nil
}

// Example Tree
//       1
//     /   \
//    2     3
//   / \   / \
//  4   5 5   6
func BinaryTree(numbers []int) *Tree {
	root := numbers[0]
	tree := Tree{root, nil, nil, 0, nil}
	for _, value := range numbers[1:] {
		tree.Insert(value)
	}
	return &tree
}
