package main

import (
	"fmt"
	"math"
)

type Graph map[string]map[string]int

func NewGraph() Graph {
	return make(Graph)
}

func (g Graph) AddEdge(from, to string, weight int) {
	// Initialize the nested map if it doesn't exist yet
	if g[from] == nil {
		g[from] = make(map[string]int)
	}
	g[from][to] = weight
}

func (g Graph) ShortestPath(start string) map[string]int {
	// Initialize the shortest paths with a high value
	shortest := make(map[string]int)
	for vertex := range g {
		shortest[vertex] = math.MaxInt32
	}
	shortest[start] = 0

	unvisited := make(map[string]bool)
	for vertex := range g {
		unvisited[vertex] = true
	}

	for len(unvisited) > 0 {
		// Find the vertex with the shortest tentative distance
		var next string
		for vertex := range unvisited {
			if next == "" || shortest[vertex] < shortest[next] {
				next = vertex
			}
		}

		// Update the shortest paths to the neighboring vertices
		for neighbor, weight := range g[next] {
			tentative := shortest[next] + weight
			if tentative < shortest[neighbor] {
				shortest[neighbor] = tentative
			}
		}

		// Mark the vertex as visited
		delete(unvisited, next)
	}

	return shortest
}

func main() {
	g := NewGraph()
	g.AddEdge("a", "b", 7)
	g.AddEdge("a", "c", 9)
	g.AddEdge("a", "f", 14)
	g.AddEdge("b", "c", 10)
	g.AddEdge("b", "d", 15)
	g.AddEdge("c", "d", 11)
	g.AddEdge("c", "f", 2)
	g.AddEdge("d", "e", 6)
	g.AddEdge("e", "f", 9)

	shortest := g.ShortestPath("a")
	for vertex, distance := range shortest {
		fmt.Printf("Shortest distance from a to %s: %d\n", vertex, distance)
	}
}

