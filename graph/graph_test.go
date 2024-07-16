package graph

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

type graphTest struct {
	vertecies     []string
	edges         []edgeTest
	expectedBFS   []string
	expectedDFS   []string
	BFSStartPoint string
	DFSStartPoint string
}

type edgeTest struct {
	from string
	to   string
}

func TestBFS_DFS(t *testing.T) {
	tests := []graphTest{
		{
			vertecies:     []string{"a", "b", "c", "d", "e", "f"},
			BFSStartPoint: "a",
			DFSStartPoint: "a",
			edges: []edgeTest{
				{from: "a", to: "b"},
				{from: "b", to: "e"},
				{from: "e", to: "f"},
				{from: "a", to: "c"},
				{from: "c", to: "d"},
			},
			expectedBFS: []string{"a", "b", "c", "e", "d", "f"},
			expectedDFS: []string{"a", "b", "e", "f", "c", "d"},
		},
	}

	for _, c := range tests {
		g := NewGraph()
		for _, v := range c.vertecies {
			g.AddVertex(v)
		}
		for _, e := range c.edges {
			g.AddEdge(e.from, e.to, 0)
		}
		resultBFS := make(chan string)
		resultDFS := make(chan string)

		go func() {
			g.BFS(c.BFSStartPoint, resultBFS)
		}()

		go func() {
			g.DFS(c.DFSStartPoint, resultDFS)
		}()
		bfs := []string{}
		dfs := []string{}
		for i := range resultBFS {
			bfs = append(bfs, i)
		}
		for i := range resultDFS {
			dfs = append(dfs, i)
		}
		if !slices.Equal(bfs, c.expectedBFS) {
			t.Errorf("expected %v but got %v", c.expectedBFS, bfs)
		}
		if !slices.Equal(dfs, c.expectedDFS) {
			t.Errorf("expected %v but got %v", c.expectedDFS, dfs)
		}

	}
}

func TestGet(t *testing.T) {
	assert := assert.New(t)
	g := NewGraph()
	g.AddVertex("a")
	g.AddVertex("b")
	g.AddVertex("d")
	g.AddEdge("a", "b", 1)
	r := g.Get("a")
	assert.Equal([]Edge{{To: "b", Weight: 1}}, r)
	r = g.Get("c")
	assert.Nil(r)
	r = g.Get("d")
	assert.NotNil(r)
}
