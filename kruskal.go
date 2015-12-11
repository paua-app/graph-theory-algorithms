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

	//Output of the edge map
	fmt.Println("\n\nEdge map:\n")
	for _, e := range g.edgeMap {
		fmt.Println(e)
	}

	// Start Kruskal algorithm
	fmt.Println("\n\nStart kruskal:\n")
	A := make([]edge, 0)
	// iterates over the list of sorted keys (e)
	for e, i := getBounds(g.edgeMap), 0; i < len(e); i++ {
		fmt.Println(i, " - ", e[i])
		for _, e := range g.edgeMap[e[i]] {
			fmt.Println("    ", e)
			if findSet(&e.x) != findSet(&e.y) {
				A = append(A, e)
				union(&e.x, &e.y)
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
func getBounds(edges map[int][]edge) []int {
	output := make([]int, 0)

	// get all keys
	for i, _ := range edges {
		output = append(output, i)
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
