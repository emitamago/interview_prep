package main

import (
	"fmt"
)

// Define vertex
type Vertex struct {
	key      int
	adjacent []*Vertex
}

// Define Graph
type Graph struct {
	vertices []*Vertex
}

// Define AddVertex to graph
func (g *Graph) AddVertex(k int) {
	if !contain(g.vertices, k) {
		g.vertices = append(g.vertices, &Vertex{key: k})
	} else {
		fmt.Printf("%v is already in the graph\n", k)
	}
}

// Define AddEdge to add edge
func (g *Graph) AddEdge(from, to int) {
	fromVertex := g.GetVertex(from)
	toVertex := g.GetVertex(to)
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("invalid edge (%v --> %v)", from, to)
		fmt.Println(err.Error())
	} else if contain(fromVertex.adjacent, to) {
		err := fmt.Errorf("edge already exist (%v --> %v)", from, to)
		fmt.Println(err.Error())
	} else {
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	}

}

// check if key exist in graph already
func contain(list []*Vertex, k int) bool {
	for _, v := range list {
		if v.key == k {
			return true
		}
	}
	return false
}

// Define GetVertex to return vertex with key
func (g *Graph) GetVertex(key int) *Vertex {
	for _, v := range g.vertices {
		if v.key == key {
			return v
		}
	}
	return nil
}

// Print all contents of graph
func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("Vertex %v\n", v.key)
		for _, v := range v.adjacent {
			fmt.Printf("%v", v.key)
		}
	}
}

func main() {
	test := &Graph{}

	for i := 0; i < 5; i++ {
		test.AddVertex(i)
	}

	// Add exiting vertex
	test.AddVertex(2)

	// invalid from and to
	test.AddEdge(10, 3)

	// valid edge
	test.AddEdge(1, 4)

	// edge already exist
	test.AddEdge(1, 4)
}
