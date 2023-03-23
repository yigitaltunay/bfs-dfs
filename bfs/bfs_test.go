package bfs

import (
	"testing"
)

func TestBFSV1(t *testing.T) {
	nodeA := &Node{value: "A"}
	nodeB := &Node{value: "B"}
	nodeC := &Node{value: "C"}
	nodeD := &Node{value: "D"}
	nodeE := &Node{value: "E"}
	nodeF := &Node{value: "F"}
	nodeG := &Node{value: "G"}

	nodeA.edges = []*Node{nodeB, nodeC}
	nodeB.edges = []*Node{nodeD, nodeE}
	nodeC.edges = []*Node{nodeE}
	nodeD.edges = []*Node{}
	nodeE.edges = []*Node{nodeF}
	nodeF.edges = []*Node{nodeG}
	nodeG.edges = []*Node{}

	result := bfsV1(nodeA)

	expected := []string{"A", "B", "C", "D", "E", "F", "G"}
	for i, val := range expected {
		if result[i] != val {
			t.Errorf("Expected %s at index %d, but got %s", val, i, result[i])
		}
	}
}

func TestBFSV2(t *testing.T) {
	nodeA := &Node{value: "A"}
	nodeB := &Node{value: "B"}
	nodeC := &Node{value: "C"}
	nodeD := &Node{value: "D"}
	nodeE := &Node{value: "E"}
	nodeF := &Node{value: "F"}
	nodeG := &Node{value: "G"}

	nodeA.edges = []*Node{nodeB, nodeC}
	nodeB.edges = []*Node{nodeD, nodeE}
	nodeC.edges = []*Node{nodeE}
	nodeD.edges = []*Node{}
	nodeE.edges = []*Node{nodeF}
	nodeF.edges = []*Node{nodeG}
	nodeG.edges = []*Node{}

	result := bfsV2(nodeA)

	expected := []string{"A", "B", "C", "D", "E", "F", "G"}
	for i, val := range expected {
		if result[i] != val {
			t.Errorf("Expected %s at index %d, but got %s", val, i, result[i])
		}
	}
}
