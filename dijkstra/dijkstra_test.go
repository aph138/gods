package dijkstra

import (
	"slices"
	"testing"

	"github.com/aph138/gods/graph"
)

type edgeTest struct {
	from   string
	to     string
	weight int
}
type dijTest struct {
	vertices   []string
	edges      []edgeTest
	start, end string
	expected   []string
	err        error
}

func TestMain(t *testing.T) {
	tests := []dijTest{
		{
			vertices: []string{"a", "b", "c", "d", "e"},
			start:    "a",
			end:      "e",
			edges: []edgeTest{
				{from: "a", to: "b", weight: 1},
				{from: "b", to: "c", weight: 2},
				{from: "c", to: "e", weight: 1},
				{from: "a", to: "d", weight: 4},
				{from: "d", to: "e", weight: 1},
			},
			expected: []string{"a", "b", "c", "e"},
			err:      nil,
		},
		{
			vertices: []string{"a"},
			start:    "c",
			err:      ErrStartNotFound,
		},
		{
			vertices: []string{"a"},
			start:    "a",
			end:      "c",
			err:      ErrEndNotFound,
		},
		{
			vertices: []string{"a", "b", "c", "d", "e"},
			start:    "a",
			end:      "e",
			edges: []edgeTest{
				{from: "a", to: "b", weight: 1},
				{from: "b", to: "c", weight: 2},
				{from: "a", to: "d", weight: 4},
			},
			err: ErrNoWay,
		},
	}
	for _, c := range tests {
		g := graph.NewGraph()
		for _, v := range c.vertices {
			g.AddVertex(v)
		}
		for _, e := range c.edges {
			g.AddEdge(e.from, e.to, e.weight)
		}
		result, err := Dijkstra(g, c.start, c.end)
		if err != c.err {
			t.Errorf("expected %v but got %v", c.err, err)
		}
		if slices.Compare(result, c.expected) != 0 {
			t.Errorf("expected %v but got %v", c.expected, result)
		}

	}
}
