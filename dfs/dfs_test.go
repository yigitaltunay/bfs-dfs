package dfs

import "testing"

func TestDFS(t *testing.T) {
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

	visited := make(map[*Node]bool)
	result := []string{}
	result = dfs(nodeA, visited, result)

	expected := []string{"A", "B", "D", "E", "F", "G", "C"}
	for i, val := range expected {
		if result[i] != val {
			t.Errorf("Expected %s at index %d, but got %s", val, i, result[i])
		}
	}
}
