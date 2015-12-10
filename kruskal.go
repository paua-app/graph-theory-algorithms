package main

import (
	"fmt"
)

func main() {
	g := newGraph()

	//TODO read graphs from file

	g.addVertex("a", 0)
	g.addVertex("b", 0)
	g.addVertex("c", 0)
	g.addVertex("d", 0)
	g.addVertex("e", 0)

	g.addEdge("a", "b", 2)
	g.addEdge("a", "e", 1)
	g.addEdge("b", "e", 42)
	g.addEdge("b", "d", 1)
	g.addEdge("b", "c", 3)
	g.addEdge("c", "d", 7)
	g.addEdge("e", "d", 2)

	//TODO better output + output of minimal spanning tree

	fmt.Println(g.vertexMap)
	fmt.Println(g.edgeMap)
}
