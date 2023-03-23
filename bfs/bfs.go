package bfs

import (
	"container/list"
)

type Node struct {
	value string
	edges []*Node
}

func bfsV1(start *Node) []string {
	visited := make(map[*Node]bool)
	queue := list.New()
	queue.PushBack(start)
	visited[start] = true
	result := []string{}

	for queue.Len() != 0 {
		current := queue.Front().Value.(*Node)
		queue.Remove(queue.Front())
		result = append(result, current.value)

		for _, n := range current.edges {
			if !visited[n] {
				visited[n] = true
				queue.PushBack(n)
			}
		}
	}
	return result
}

func bfsV2(start *Node) []string {
	visited := make(map[*Node]bool)
	queue := []*Node{start}
	visited[start] = true
	result := []string{}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		result = append(result, current.value)
		for _, n := range current.edges {
			if !visited[n] {
				visited[n] = true
				queue = append(queue, n)
			}
		}
	}
	return result
}
