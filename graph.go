package main

type graph struct {
	vertexMap map[string]*vertex
	edgeMap   *[]edge
}

func newGraph() graph {
	vertexMap := make(map[string]*vertex, 0)
	edgeMap := make([]edge, 0)
	return graph{vertexMap: vertexMap, edgeMap: &edgeMap}
}

// Creates a vertex and adds it to the graph.
// When the vertex already exists, it'll be overwritten.
func (g graph) addVertex(name string) {
	v := newVertex(name)
	g.vertexMap[name] = &v
}

// Creates an edge to the graph
func (g *graph) addEdgeWeight(v1, v2 string, weight int) {
	edge := newEdgeWeight(*g.vertexMap[v1], *g.vertexMap[v2], weight)
	g.addEdge(edge)
}

func (g graph) addEdgeFlow(v1, v2 string, flow, capacity int) {
	edge := newEdgeFlow(*g.vertexMap[v1], *g.vertexMap[v2], flow, capacity)
	g.addEdge(edge)
}

func (g *graph) addEdge(edge edge) {
	contains := false
	for _, e := range *g.edgeMap {
		if e.equals(edge) {
			contains = true
			break
		}
	}
	if !contains {
		// add reachables:
		x := g.vertexMap[edge.x.name]
		y := g.vertexMap[edge.y.name]
		reachables := append(*x.reachables, *y)
		x.reachables = &reachables

		edgeMap := append(*g.edgeMap, edge)
		g.edgeMap = &edgeMap
	}
}

func (g graph) getEdgesWithWeight(w int) []edge {
	output := make([]edge, 0)
	for _, e := range *g.edgeMap {
		if e.weight == w {
			output = append(output, e)
		}
	}
	return output
}
