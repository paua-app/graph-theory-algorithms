package main

type graph struct {
	vertexMap map[string]vertex
	edgeMap   map[int][]edge
}

func newGraph() graph {
	//TODO check for invalid values in v and e
	vertexMap := make(map[string]vertex, 1)
	edgeMap := make(map[int][]edge, 1)
	return graph{vertexMap: vertexMap, edgeMap: edgeMap}
}

// Creates a vertex and adds it to the graph.
// When the vertex already exists, it'll be overwritten.
func (g graph) addVertex(name string, color int) {
	v := vertex{name: name, color: color}
	g.vertexMap[name] = v
}

// Creates an edge to the graph
func (g graph) addEdge(v1, v2 string, weight int) {
	e := newEdge(g.vertexMap[v1], g.vertexMap[v2])

	if _, doesExist := g.edgeMap[weight]; doesExist {
		g.edgeMap[weight] = append(g.edgeMap[weight], e)
	} else {
		g.edgeMap[weight] = []edge{e}
	}
}
