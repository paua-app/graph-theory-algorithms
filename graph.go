package main

type graph struct {
	vertexMap map[string]vertex
	edgeMap   []edge
}

func newGraph() graph {
	//TODO check for invalid values in v and e
	vertexMap := make(map[string]vertex, 0)
	edgeMap := make([]edge, 0)
	return graph{vertexMap: vertexMap, edgeMap: edgeMap}
}

// Creates a vertex and adds it to the graph.
// When the vertex already exists, it'll be overwritten.
func (g graph) addVertex(name string) {
	v := newVertex(name)
	g.vertexMap[name] = v
}

// Creates an edge to the graph
func (g *graph) addEdge(v1, v2 string, weight int) {
	edge := newEdge(g.vertexMap[v1], g.vertexMap[v2], weight)
	contains := false
	for _, e := range g.edgeMap {
		if e.equals(edge) {
			contains = true
			break
		}
	}
	if !contains {
		g.edgeMap = append(g.edgeMap, edge)
		e := &edge
		e.x.neighbors = append(e.x.neighbors, e.y)
		e.y.neighbors = append(e.y.neighbors, e.x)
	}
}

func (g graph) getEdgesWithWeight(w int) []edge {
	output := make([]edge, 0)
	for _, e := range g.edgeMap {
		if e.weight == w {
			output = append(output, e)
		}
	}
	return output
}
