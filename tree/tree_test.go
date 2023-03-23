package tree

import (
	"testing"
)

func TestBinaryTree(t *testing.T) {
	numbers := []int{5, 2, 8, 1, 3, 7, 9}
	tree := BinaryTree(numbers)

	if tree.Value != 5 {
		t.Errorf("Expected root value to be 5, got %d", tree.Value)
	}

	if tree.Left.Value != 2 {
		t.Errorf("Expected left child of root to be 2, got %d", tree.Left.Value)
	}

	if tree.Right.Value != 8 {
		t.Errorf("Expected right child of root to be 8, got %d", tree.Right.Value)
	}

	repeat, _, found := tree.Find(2)
	if !found {
		t.Errorf("Expected to find value 2 in the tree")
	}

	if repeat != 1 {
		t.Errorf("Expected repeat count of 2 to be 1, got %d", repeat)
	}

	err := tree.Delete(2)
	if err != nil {
		t.Errorf("Expected no error while deleting value 2, got %v", err)
	}

	_, _, found = tree.Find(2)
	if found {
		t.Errorf("Expected value 2 to be deleted from the tree")
	}

	minimum := tree.Minimum()
	if minimum != 1 {
		t.Errorf("Expected minimum value in the tree to be 1, got %d", minimum)
	}

	maximum := tree.Maximum()
	if maximum != 9 {
		t.Errorf("Expected maximum value in the tree to be 9, got %d", maximum)
	}
}
