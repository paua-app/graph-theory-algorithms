package main

type vertex struct {
	name        string
	color, rank int
	parent      *vertex
}

func newVertex(name string, color int) vertex {
	//TODO check for invalid values
	v := vertex{name: name, color: color}
	makeSet(&v)
	return v
}

// Checks if two vertices are equal. They are equal when their names are egual.
func (v vertex) equals(v2 vertex) bool {
	return v.name == v2.name
}

// Initializes a set by adding the vertex as first element and as representing
// vertex of the set.
func makeSet(x *vertex) {
	x.parent = x
	x.rank = 0
}

// Merges to sets.
func union(x, y *vertex) {
	link(findSet(x), findSet(y))
}

// Merges to sets and updates the ranks of the vertices.
func link(x, y *vertex) {
	if x.rank > y.rank {
		y.parent = x
	} else {
		x.parent = y
		if x.rank == y.rank {
			y.rank++
		}
	}
}

// Gets the representation vertex of the set (= parent).
// This also compresses the path of the set
func findSet(x *vertex) *vertex {
	if x.parent != x {
		x.parent = findSet(x.parent)
	}
	return x.parent
}
