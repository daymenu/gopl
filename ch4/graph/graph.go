package graph

var graph = make(map[string]map[string]bool)

func AddEdge(from, to string) {
	edges := graph[from]

	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}
func HasEdge(from, to string) bool {
	return graph[from][to]
}
