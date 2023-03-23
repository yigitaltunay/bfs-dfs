package dfs

type Node struct {
	value string
	edges []*Node
}

func dfs(node *Node, visited map[*Node]bool, result []string) []string {
	visited[node] = true
	result = append(result, node.value)
	for _, n := range node.edges {
		if !visited[n] {
			result = dfs(n, visited, result)
		}
	}
	return result
}
