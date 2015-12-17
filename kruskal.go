package main

import (
	"fmt"
)

func main() {
	g := newGraph()

	//TODO read graphs from file

	g.addVertex("a")
	g.addVertex("b")
	g.addVertex("c")
	g.addVertex("d")
	g.addVertex("e")

	g.addEdge("a", "b", 2)
	g.addEdge("a", "e", 1)
	g.addEdge("b", "e", 42)
	g.addEdge("b", "d", 1)
	g.addEdge("b", "c", 3)
	g.addEdge("c", "d", 7)
	g.addEdge("e", "d", 2)

	//TODO better output + output of minimal spanning tree

	// Start Kruskal algorithm
	fmt.Println("\n\nStart kruskal:\n")
	A := make([]edge, 0)
	// iterates over the list of sorted keys (e)
	for e, i := getBounds(g.edgeMap), 0; i < len(e); i++ {
		fmt.Println(i, " - ", e[i])
		for _, e := range g.getEdgesWithWeight(e[i]) {
			fmt.Println("    ", e)
			if (&e.x).findSet() != (&e.y).findSet() {
				A = append(A, e)
				(&e.x).union(&e.y)
			}
		}
	}
	fmt.Println("\n\nTree:\n")

	//Output of the tree
	for _, e := range A {
		fmt.Println(e)
	}
}

// Gets a sorted list of all keys of the given map.
func getBounds(edges []edge) []int {
	output := make([]int, 0)

	// get all keys
	for _, i := range edges {
		output = append(output, i.weight)
	}

	// sort list of keys
	// this uses bubble sort because it's easy
	fmt.Println("pre: ", output)
	for n := len(output); n > 1; n-- {
		for i := 0; i < n-1; i++ {
			if output[i] > output[i+1] {
				temp := output[i]
				output[i] = output[i+1]
				output[i+1] = temp
			}
		}
	}
	fmt.Println("after: ", output)
	return output
}
