package main

type vertex struct {
	name             string
	rank, key, color int
	parent           *vertex
	reachables       *[]vertex
}

func newVertex(name string) vertex {
	reachables := make([]vertex, 0)
	v := vertex{name: name, reachables: &reachables}
	(&v).makeSet()
	return v
}

func (v *vertex) setKey(key int) {
	v.key = key
}

func (v *vertex) setRank(rank int) {
	v.rank = rank
}

func (v *vertex) setColor(color int) {
	v.color = color
}

func (v *vertex) setParent(parent *vertex) {
	v.parent = parent
}

// Checks if two vertices are equal. They are equal when their names are egual.
func (v vertex) equals(v2 vertex) bool {
	return v.name == v2.name
}

// Initializes a set by adding the vertex as first element and as representing
// vertex of the set.
func (v *vertex) makeSet() {
	v.parent = v
	v.rank = 0
}

// Merges to sets.
func (v *vertex) union(y *vertex) {
	v.findSet().link(y.findSet())
}

// Merges to sets and updates the ranks of the vertices.
func (v *vertex) link(y *vertex) {
	if v.rank > y.rank {
		y.parent = v
	} else {
		v.parent = y
		if v.rank == y.rank {
			y.rank++
		}
	}
}

// Gets the representation vertex of the set (= parent).
// This also compresses the path of the set
func (v *vertex) findSet() *vertex {
	if v.parent != v {
		v.parent = v.parent.findSet()
	}
	return v.parent
}
