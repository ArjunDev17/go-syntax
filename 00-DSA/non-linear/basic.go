package main

import (
	"fmt"
)

type Graph struct {
	vertices map[string][]string
}

func NewGraph() *Graph {
	return &Graph{
		vertices: make(map[string][]string),
	}
}

func (g *Graph) AddEdge(from, to string) {
	g.vertices[from] = append(g.vertices[from], to)
}

func (g *Graph) Print() {
	for node, edges := range g.vertices {
		fmt.Printf("%s -> %v\n", node, edges)
	}
}

func main() {
	graph := NewGraph()

	graph.AddEdge("A", "B")
	graph.AddEdge("A", "C")
	graph.AddEdge("B", "D")
	graph.AddEdge("C", "D")

	graph.Print()
}
