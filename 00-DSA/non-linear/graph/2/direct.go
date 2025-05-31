package main

import (
	"fmt"
)

type Graph struct {
	vertices map[string][]string
}

// Create a new graph
func NewGraph() *Graph {
	return &Graph{
		vertices: make(map[string][]string),
	}
}

// Add a directed edge from one vertex to another
func (g *Graph) AddEdge(from, to string) {
	g.vertices[from] = append(g.vertices[from], to)
}

// Print the graph's adjacency list
func (g *Graph) Print() {
	fmt.Println("Graph:")
	for vertex, edges := range g.vertices {
		fmt.Printf("%s -> %v\n", vertex, edges)
	}
}

// Depth-First Search (recursive)
func (g *Graph) DFS(start string, visited map[string]bool) {
	if visited[start] {
		return
	}
	visited[start] = true
	fmt.Print(start, " ")

	for _, neighbor := range g.vertices[start] {
		g.DFS(neighbor, visited)
	}
}

// Breadth-First Search (using a queue)
func (g *Graph) BFS(start string) {
	visited := make(map[string]bool)
	queue := []string{start}
	visited[start] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		fmt.Print(current, " ")

		for _, neighbor := range g.vertices[current] {
			if !visited[neighbor] {
				queue = append(queue, neighbor)
				visited[neighbor] = true
			}
		}
	}
}

func main() {
	graph := NewGraph()

	// Add directed edges
	graph.AddEdge("A", "B")
	graph.AddEdge("A", "C")
	graph.AddEdge("B", "D")
	graph.AddEdge("C", "D")
	graph.AddEdge("D", "E")

	graph.Print()

	fmt.Print("\nDFS starting from A: ")
	graph.DFS("A", make(map[string]bool))

	fmt.Print("\nBFS starting from A: ")
	graph.BFS("A")

	fmt.Println()
}
