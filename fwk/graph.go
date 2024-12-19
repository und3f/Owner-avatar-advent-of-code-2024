package fwk

import (
	"slices"

	"github.com/und3f/aoc/2024/fwk/twoD"
)

type Graph interface {
	Len() int
	Adj(v int) []int
}

type BFS struct {
	graph Graph
	to    []int
	start int
}

func NewBFS(graph Graph, v int) *BFS {
	to := make([]int, graph.Len())
	bfs := &BFS{graph, to, v}
	bfs.bfs()

	return bfs
}

func (bfs *BFS) bfs() {
	visited := make([]bool, bfs.graph.Len())
	for i := range bfs.to {
		bfs.to[i] = -1
	}
	bfs.to[bfs.start] = bfs.start

	heap := []int{bfs.start}
	for len(heap) > 0 {
		v := heap[0]
		heap = heap[1:]

		if visited[v] {
			continue
		}
		visited[v] = true

		for _, w := range bfs.graph.Adj(v) {
			if visited[w] {
				continue
			}

			if bfs.to[w] == -1 {
				bfs.to[w] = v
			}
			heap = append(heap, w)
		}
	}
}

func (bfs *BFS) PathTo(target int) []int {
	var path []int

	for v := target; v != bfs.start; v = bfs.to[v] {
		path = append(path, v)
		if v == -1 {
			return nil
		}
	}
	path = append(path, bfs.start)
	slices.Reverse(path)
	return path
}

type GraphMask struct {
	graph Graph
	mask  []bool
}

func (g *GraphMask) Len() int {
	return g.graph.Len()
}

func (g *GraphMask) Adj(v int) []int {
	adj := g.graph.Adj(v)
	for i := len(adj) - 1; i >= 0; i-- {
		if g.mask[adj[i]] {
			newAdj := adj[:i]
			if i+1 < len(adj) {
				newAdj = append(adj, adj[i+1:]...)
			}
			adj = newAdj
		}
	}

	return adj
}

type GridGraph struct {
	flattenGrid    []rune
	n              int
	width          int
	fourDirections []int
}

func NewGridGraph(grid [][]rune) Graph {
	w := 0
	n := 0
	if len(grid) > 0 {
		w = len(grid[0])
		n = len(grid) * w
	}

	fourDirections := make([]int, 4)
	for i, direction := range [][]int{
		twoD.DirectionEast,
		twoD.DirectionWest,
		twoD.DirectionNorth,
		twoD.DirectionSouth,
	} {
		fourDirections[i] = HashVect(grid, direction)
	}

	flattenGrid := make([]rune, n)
	for i, row := range grid {
		for j, v := range row {
			flattenGrid[i*w+j] = v
		}
	}

	return &GridGraph{
		flattenGrid, n, w, fourDirections,
	}
}

func (g GridGraph) Len() int {
	return g.n
}

func (g GridGraph) Adj(v int) []int {
	var adj []int
	for _, vec := range g.fourDirections {
		w := v + vec
		if w < 0 || w >= g.n {
			continue
		}
		switch vec {
		case -1:
			if v%g.width == 0 {
				continue
			}
		case 1:
			if v%g.width == g.width-1 {
				continue
			}
		}

		if g.flattenGrid[w] == '#' {
			continue
		}

		adj = append(adj, w)
	}

	return adj
}
