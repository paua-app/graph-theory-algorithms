package main

type edge struct {
	x, y vertex
}

// Creates a new edge between x and y with a specifig wgith.
// The weight must be greater than 0
func newEdge(x, y vertex) edge {
	//TODO check if vertices are valid (e.g. x != y)
	return edge{x: x, y: y}
}

// Checks if two edges are the same meaning if they contain the same vertices.
func (e edge) equals(e2 edge) bool {
	return e.x.equals(e2.y) && e.y.equals(e2.x) ||
		e.x.equals(e2.x) && e.y.equals(e2.y)
}
