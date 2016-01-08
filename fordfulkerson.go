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
	network.g.addEdgeFlow("w", "u", 0, 9)
	network.g.addEdgeFlow("x", "w", 0, 7)
	network.g.addEdgeFlow("w", "t", 0, 20)
	network.g.addEdgeFlow("x", "t", 0, 4)

	getPath(&network.g.vertexMap, *network.g.vertexMap["s"], *network.g.vertexMap["t"])

	fmt.Println("\n\n-------- << == >> --------")
	fmt.Println("The path is: ")

	node := network.g.vertexMap["t"]
	output := ""

	for node.name != "s" {
		output = "  -->  " + node.name + output
		node = node.parent
	}

	output = node.parent.name + output
	fmt.Println(output)
}

// returns a path through the graph
func getPath(vertexMap *map[string]*vertex, start, goal vertex) {
	for _, v := range *vertexMap {
		v.setColor(0)
	}

	// cormen suggests a loop but this also works like a charm
	DFSVisit((*vertexMap)["s"], (*vertexMap)["t"], vertexMap)
}

func DFSVisit(u, goal *vertex, vertexMap *map[string]*vertex) {
	if u.equals(*goal) {
		fmt.Println("GOAL: ", u.name, " - ", goal.name)
		return
	}

	u.color = 1
	fmt.Println(u.name, ": ", len(*u.reachables))

	for _, v1 := range *u.reachables {
		v := (*vertexMap)[v1.name]
		fmt.Println("Reachable: ", u.name, " - ", v.name)

		if v.color == 0 {
			fmt.Println("add parent: ", v.name, ".parent = ", u.name)
			v.setParent(u)
			DFSVisit(v, goal, vertexMap)
		}
	}
	return
}
