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
func (g *Graph) AddEdge1(from, to string) {
	g.vertices[from] = append(g.vertices[from], to)
	g.vertices[to] = append(g.vertices[to], from)
}
func (g *Graph) DFS(start string, visited map[string]bool) {
	if visited[start] {
		return
	}
	visited[start] = true
	fmt.Println(start)

	for _, neighbor := range g.vertices[start] {
		g.DFS(neighbor, visited)
	}
}

func main() {
	graph := NewGraph()

	graph.AddEdge("A", "B")
	graph.AddEdge("A", "C")
	graph.AddEdge("B", "D")
	graph.AddEdge("C", "D")

	graph.Print()
	visited := make(map[string]bool)
	graph.DFS("A", visited)
}
