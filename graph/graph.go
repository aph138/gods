package graph

import (
	"fmt"
	"log"

	"github.com/aph138/gods/queue"
)

type Edge struct {
	To     string
	Weight int
}

type Graph struct {
	vertices map[string][]Edge
	directed bool
	verbose  bool
}

type configFunc func(*Graph)

func WithDirection() configFunc {
	return func(g *Graph) {
		g.directed = true
	}
}
func WithVerbosity() configFunc {
	return func(g *Graph) {
		g.verbose = true
	}
}

// NewGraph creates a new graph. You can pass config function to it.
// config functions:
//   - WithVerbosity() -> BFS and DFS will print any node that they visit
//   - WithDirection() -> AddEdge will add the edge for only the "from" vertex. Otherwise
//     it will add edge for both "from" and "to" vertices.
func NewGraph(configs ...configFunc) *Graph {
	g := &Graph{
		vertices: make(map[string][]Edge),
	}
	for _, i := range configs {
		i(g)
	}
	return g
}

func (g *Graph) AddVertex(name string) {
	if _, exist := g.vertices[name]; !exist {
		g.vertices[name] = []Edge{}
	}
}

func (g *Graph) AddEdge(from, to string, weight int) {
	if _, exist := g.vertices[from]; exist {
		g.vertices[from] = append(g.vertices[from], Edge{To: to, Weight: weight})
	}
	if !g.directed {
		_, exist := g.vertices[to]
		if !exist {
			g.AddVertex(to)
		}
		g.vertices[to] = append(g.vertices[to], Edge{To: from, Weight: weight})
	}
}

// Get returns edges of the given vertex
func (g *Graph) Get(vertex string) []Edge {
	return g.vertices[vertex]

}

// Vertices returns list of vertices of the graph
func (g *Graph) Vertices() []string {
	l := make([]string, len(g.vertices))

	for k := range g.vertices {
		l = append(l, k)
	}

	return l
}

// Len returns number of vertices
func (g *Graph) Len() int {
	return len(g.vertices)
}
func (g *Graph) Display() {
	for i, k := range g.vertices {
		fmt.Printf("%s ->\n", i)
		for _, edge := range k {
			fmt.Printf("...%s (%d)\n", edge.To, edge.Weight)
		}
	}
}

// DFS performs a Depth-First Search on the graph from the given start point.
// It sends the nodes that are visited in a depth-first manner to the given channel.
// The method will close the channel upon the completion of the traversal.
func (g *Graph) DFS(src string, result chan string) {
	g.dfsHelder(src, make(map[string]bool), result)
	close(result)
}
func (g *Graph) dfsHelder(src string, visited map[string]bool, result chan string) {
	if edges, ok := g.vertices[src]; ok {
		if visited[src] {
			return
		}
		visited[src] = true
		result <- src
		if g.verbose {
			log.Printf("visited %s", src)
		}
		for _, e := range edges {
			g.dfsHelder(e.To, visited, result)
		}
	}
}

// BFS performs a Breadth-First Search on the graph from the given start point.
// It sends the nodes that are visited in a breadth-first manner to the given channel.
// The method will close the channel upon the completion of the traversal.
func (g *Graph) BFS(start string, reuslt chan string) {
	if _, ok := g.vertices[start]; !ok {
		return
	}
	queue := queue.NewQueue(start)
	visited := make(map[string]bool)
	visited[start] = true
	reuslt <- start
	for queue.Len() > 0 {
		p, _ := queue.Dequeue()
		for _, edge := range g.vertices[p] {
			if !visited[edge.To] {
				reuslt <- edge.To
				if g.verbose {
					fmt.Printf("visited %s\n", edge.To)
				}
				visited[edge.To] = true
				queue.Enqueue(edge.To)
			}
		}
	}
	close(reuslt)
}
