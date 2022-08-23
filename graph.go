package main

type Vertex = string

type Edge struct {
	from Vertex
	to   Vertex
}

func NewEdge(from Vertex, to Vertex) Edge {
	return Edge{from: from, to: to}
}

type Graph struct {
	AdjLists map[Vertex][]Vertex
}

func NewGraph(edges []Edge) Graph {
	var adjLists = make(map[Vertex][]Vertex)
	for _, e := range edges {
		if adjList, present := adjLists[e.from]; present {
			adjLists[e.from] = append(adjList, e.to)
		} else {
			adjLists[e.from] = []Vertex{e.to}
		}
		if _, present := adjLists[e.to]; !present {
			adjLists[e.to] = make([]Vertex, 0)
		}
	}
	return Graph{AdjLists: adjLists}
}

func TopologicalSort(graph Graph) []Vertex {

	var nbVertices = len(graph.AdjLists)
	var time = 0
	var sorted = make([]Vertex, nbVertices)
	var discoveredVertices = make(map[Vertex]bool, nbVertices)

	var explore func([]Vertex)
	explore = func(vertices []Vertex) {
		for _, u := range vertices {
			if !discoveredVertices[u] {
				explore(graph.AdjLists[u])
				time += 1
				sorted[nbVertices-time] = u
				discoveredVertices[u] = true
			}
		}
	}

	var vertices = make([]Vertex, 0, nbVertices)
	for u := range graph.AdjLists {
		vertices = append(vertices, u)
	}
	explore(vertices)
	return sorted
}
