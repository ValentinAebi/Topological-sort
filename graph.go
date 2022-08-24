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

func (graph *Graph) getVertices() []Vertex {
	var nbVertices = len(graph.AdjLists)
	var vertices = make([]Vertex, 0, nbVertices)
	for u := range graph.AdjLists {
		vertices = append(vertices, u)
	}
	return vertices
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
				discoveredVertices[u] = true
				explore(graph.AdjLists[u])
				time += 1
				sorted[nbVertices-time] = u
			}
		}
	}

	explore(graph.getVertices())
	return sorted
}

func FindCycle(graph Graph) []Vertex {

	var nbVertices = len(graph.AdjLists)
	var currentPathVertices = make(map[Vertex]int, nbVertices)
	var depth = 0

	// Return:
	//
	// - stack: the stack at the time of cycle detection: Vertex -> depth
	//
	// - firstRepeated: the vertex in the stack at which the cycle begins
	var explore func([]Vertex) (stack map[Vertex]int, firstRepeated Vertex)
	explore = func(vertices []Vertex) (stack map[Vertex]int, firstRepeated Vertex) {
		for _, u := range vertices {
			if _, pr := currentPathVertices[u]; pr {
				stack = currentPathVertices
				firstRepeated = u
				return
			} else {
				currentPathVertices[u] = depth
				depth += 1
				stack, firstRepeated = explore(graph.AdjLists[u])
				depth -= 1
				if stack != nil {
					return
				}
				delete(currentPathVertices, u)
			}
		}
		return nil, ""
	}

	var stack, firstRepeated = explore(graph.getVertices())

	if stack == nil {
		return nil
	} else {
		// build reverse map: depth -> vertex
		var reversedStackMap = make(map[int]Vertex)
		var depthOfFirstRepeated int
		for vertex, depth := range stack {
			reversedStackMap[depth] = vertex
			if vertex == firstRepeated {
				depthOfFirstRepeated = depth
			}
		}
		var cycleLen = len(reversedStackMap) - depthOfFirstRepeated + 1
		var cycleSlice = make([]Vertex, cycleLen)
		for i := 0; i < cycleLen-1; i++ {
			cycleSlice[i] = reversedStackMap[i+depthOfFirstRepeated]
		}
		cycleSlice[cycleLen-1] = firstRepeated
		return cycleSlice
	}

}
