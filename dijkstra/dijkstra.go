package dijkstra

import (
	"errors"
	"math"
	"slices"

	"github.com/aph138/gods/graph"
)

var ErrStartNotFound = errors.New("starting point doesn't exist")
var ErrEndNotFound = errors.New("ending point doesn't exist")
var ErrNoWay = errors.New("there is no way that connect start to the end")

// Dijkstra find the shortest path from starting vertex to the ending vertex.
// It returns an array of string from start to end.
// If start or end isn't in graph it will return ErrStartNotFound or ErrEndNotFound.
// If there is no path from start to end it will return ErrNoWay.
func Dijkstra(graph *graph.Graph, start, end string) ([]string, error) {
	// return error if starting or ending point does'nt exsit in graph
	if graph.Get(start) == nil {
		return nil, ErrStartNotFound
	}
	if graph.Get(end) == nil {
		return nil, ErrEndNotFound
	}

	distances := make(map[string]int)  // distances keep track of the distances from starting point
	parents := make(map[string]string) // parents keep track of the previous vertex
	checked := make(map[string]bool)   //keep track of the visited vertecies

	//initializing the distances
	for _, i := range graph.Vertices() {
		distances[i] = math.MaxInt
	}
	distances[start] = 0

	for len(checked) < graph.Len() {
		shortestDistance := math.MaxInt
		nearestVertex := ""

		// find the unchecked vertex with shortest distance
		for v, d := range distances {
			if !checked[v] && d < shortestDistance {
				shortestDistance = d
				nearestVertex = v
			}
		}
		if nearestVertex == "" {
			return nil, ErrNoWay
		}
		checked[nearestVertex] = true

		// update distances from nearestVertex
		for _, e := range graph.Get(nearestVertex) {
			newDis := e.Weight + shortestDistance
			if newDis < distances[e.To] {
				distances[e.To] = newDis
				parents[e.To] = nearestVertex
			}
		}
	}
	result := []string{}
	result = append(result, end)
	for {
		if parents[end] == "" {
			break
		}
		result = append(result, parents[end])
		end = parents[end]
	}
	slices.Reverse(result)
	return result, nil
}
