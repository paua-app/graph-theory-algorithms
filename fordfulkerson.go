package main

import (
	"fmt"
)

type flownetwork struct {
	g    graph
	s, t *edge
}

func fordfulkerson() {
	network := flownetwork{
		g: newGraph()}

	network.g.addVertex("s")
	network.g.addVertex("t")
	network.g.addVertex("u")
	network.g.addVertex("v")
	network.g.addVertex("w")
	network.g.addVertex("x")

	network.g.addEdgeFlow("s", "v", 0, 16)
	network.g.addEdgeFlow("s", "u", 0, 13)
	network.g.addEdgeFlow("u", "v", 0, 4)
	network.g.addEdgeFlow("v", "u", 0, 10)
	network.g.addEdgeFlow("u", "x", 0, 14)
	network.g.addEdgeFlow("v", "w", 0, 12)
	fmt.Println(network.g.vertexMap["v"])
	network.g.addEdgeFlow("w", "u", 0, 9)
	network.g.addEdgeFlow("x", "w", 0, 7)
	network.g.addEdgeFlow("w", "t", 0, 20)
	network.g.addEdgeFlow("x", "t", 0, 4)

	getPath(network.g.vertexMap, network.g.vertexMap["s"], network.g.vertexMap["t"])
	//	fmt.Println("START")
	//	node := network.g.vertexMap["t"]
	//	fmt.Println(node)
	//	for &node != nil && node.name != "s" {
	//		fmt.Println(node.name)
	//		node = *node.parent
	//	}
}

// returns a path through the graph
func getPath(vertexMap map[string]vertex, start, goal vertex) {
	for _, v := range vertexMap {
		v.setColor(0)
	}

	fmt.Println(len(start.reachables))
	for _, v := range start.reachables {
		if v.color == 0 {
			DFSVisit(v, goal)
		}
	}
}

func DFSVisit(u, goal vertex) {
	if u.equals(goal) {
		return
	}
	u.color = 1
	for _, v := range u.reachables {
		if v.color == 0 {
			fmt.Println("add parent: ", u, " - ", v)
			v.setParent(u)
			DFSVisit(v, goal)
		}
	}
	return
}
