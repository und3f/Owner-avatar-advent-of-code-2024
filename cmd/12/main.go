package main

import (
	"github.com/und3f/aoc/2024/fwk"
	"github.com/und3f/aoc/2024/fwk/twoD"
)

func main() {
	input := fwk.ReadInputRunesLines()

	part1(input)
	part2(input)
}

func part1(region [][]rune) {
	cc := NewConnectedClusters(region)
	perimeter := make([]uint64, cc.CountClusters())

	for i, row := range region {
		for j, _ := range row {
			p := []int{i, j}
			pV := fwk.HashVect(region, p)
			pVClusterId := cc.getId(pV)

			for _, direction := range twoD.FourDirections {
				adj := fwk.AddVec(p, direction)
				adjV := fwk.HashVect(region, adj)
				if fwk.IsVecIn2DBounds(region, adj) {
					if pVClusterId == cc.getId(adjV) {
						continue
					}
				}

				perimeter[pVClusterId]++
			}
		}
	}

	sum := uint64(0)
	for i := 0; i < cc.CountClusters(); i++ {
		area := cc.CountClusterVertices(i)
		sum += area * perimeter[i]
	}

	fwk.Solution(1, sum)
}

func part2(region [][]rune) {
	cc := NewConnectedClusters(region)

	sum := uint64(0)
	for i := 0; i < cc.CountClusters(); i++ {
		area := cc.CountClusterVertices(i)

		sc := NewSidesCalc(cc)
		sides := sc.CalcSides(i)

		sum += area * sides
	}

	fwk.Solution(2, sum)
}

type ConnectedClusters struct {
	region         []rune
	n              int
	cc             []int
	id             int
	visited        []bool
	fourDirections [4]int
	width          int
}

func NewConnectedClusters(region [][]rune) *ConnectedClusters {
	N := fwk.CalcBoardVertices(region)

	regionGraph := make([]rune, N)

	v := 0
	for _, row := range region {
		for _, value := range row {
			regionGraph[v] = value
			v++
		}
	}

	var fourDirections [4]int
	for i, direction := range twoD.FourDirections {
		fourDirections[i] = fwk.HashVect(region, direction)
	}

	dfs := &ConnectedClusters{
		region:         regionGraph,
		cc:             make([]int, N),
		visited:        make([]bool, N),
		fourDirections: fourDirections,
		n:              N,
		width:          len(region[0]),
	}

	dfs.buildCC()
	return dfs
}

func (dfs *ConnectedClusters) buildCC() []int {
	for v := 0; v < dfs.n; v++ {
		if !dfs.visited[v] {
			dfs.DFS(v)
			dfs.id++
		}
	}

	return dfs.cc
}

func (dfs *ConnectedClusters) getId(v int) int {
	return dfs.cc[v]
}

func (dfs *ConnectedClusters) DFS(v int) {
	if dfs.visited[v] {
		return
	}

	dfs.visited[v] = true
	dfs.cc[v] = dfs.id

	for _, direction := range dfs.fourDirections {
		w := v + direction

		if dfs.IsAdj(v, w) {
			dfs.DFS(w)
		}
	}
}

func (dfs *ConnectedClusters) IsConnected(v, w int) bool {
	if w < 0 || w >= dfs.n {
		return false
	}

	direction := w - v
	if (v%dfs.width == 0 && direction == -1) ||
		(v%dfs.width == dfs.width-1 && direction == 1) {
		return false
	}

	return true
}

func (dfs *ConnectedClusters) IsAdj(v, w int) bool {
	if !dfs.IsConnected(v, w) {
		return false
	}

	return dfs.region[v] == dfs.region[w]
}

func (dfs *ConnectedClusters) CountClusters() int {
	return dfs.id
}

func (dfs *ConnectedClusters) CountClusterVertices(id int) (count uint64) {
	for _, wId := range dfs.cc {
		if wId == id {
			count++
		}
	}

	return
}

type SidesCalc struct {
	dfs     *ConnectedClusters
	visited [][]bool
}

func NewSidesCalc(dfs *ConnectedClusters) *SidesCalc {
	visited := make([][]bool, dfs.n+1)
	for i := range visited {
		visited[i] = make([]bool, dfs.n+1)
	}

	return &SidesCalc{
		dfs:     dfs,
		visited: visited,
	}
}

func (sc *SidesCalc) CalcSides(clusterId int) uint64 {
	var count uint64

	outter := sc.FindFirstVertex(clusterId, 0)
	count += sc.CalcSidesOutside(outter, int(twoD.DirectionNorthI))

	for nextSearch := outter; nextSearch < sc.dfs.n; nextSearch++ {
		v := sc.FindFirstVertex(clusterId, nextSearch)
		if v < 0 {
			break
		}

		nextSearch = v

		for directionI, direction := range sc.dfs.fourDirections {
			w := v + direction

			if sc.dfs.IsAdj(v, w) {
				continue
			}

			if !sc.dfs.IsConnected(v, w) {
				w = sc.dfs.n
			}

			if sc.visited[v][w] {
				continue
			}

			count += sc.CalcSidesInside(w, (4+2+directionI)%4)
		}

	}

	return count
}

func (sc *SidesCalc) CalcSidesOutside(start, startDirection int) uint64 {
	dfs := sc.dfs

	v := start

	directionI := startDirection

	sides := uint64(0)

	for true {
		leftDirectionI := (4 + directionI - 1) % 4
		left := dfs.fourDirections[leftDirectionI]

		if w := v + left; dfs.IsAdj(v, w) {
			v = w
			directionI = leftDirectionI
			sides++
			continue
		} else {
			sc.RecordVisit(v, w)
		}

		if v == start && sides > 0 {
			for i := directionI; i != startDirection; i = (i + 1) % 4 {
				sides++
				sc.RecordVisit(v, v+sc.dfs.fourDirections[directionI])
			}
			break
		}

		straight := dfs.fourDirections[directionI]
		if w := v + straight; dfs.IsAdj(v, w) {
			v = w
			continue
		} else {
			sc.RecordVisit(v, w)
			sides++
		}

		rightDirectionI := (directionI + 1) % 4
		right := dfs.fourDirections[rightDirectionI]
		if w := v + right; dfs.IsAdj(v, w) {
			v = w
			directionI = rightDirectionI
			continue
		} else {
			sc.RecordVisit(v, w)
			sides++
		}

		// Should return
		directionI = (directionI + 2) % 4
		if w := v + dfs.fourDirections[directionI]; dfs.IsAdj(v, w) {
			v = w
		} else {
			for i := directionI; i != startDirection; i = (i + 1) % 4 {
				sides++
				sc.RecordVisit(v, v+sc.dfs.fourDirections[directionI])
			}
			break
		}
	}

	return sides
}

func (sc *SidesCalc) CalcSidesInside(start, startDirection int) uint64 {
	dfs := sc.dfs

	l := start + dfs.fourDirections[startDirection]
	clusterId := dfs.cc[l]

	v := start

	directionI := startDirection

	sides := uint64(0)

	for true {
		leftDirectionI := (4 + directionI - 1) % 4
		left := dfs.fourDirections[leftDirectionI]

		if w := v + left; dfs.IsConnected(v, w) && dfs.cc[w] != clusterId {
			v = w
			directionI = leftDirectionI
			sides++
			continue
		} else {
			sc.RecordVisit(w, v)
		}

		if v == start && sides > 0 {
			for i := directionI; i != startDirection; i = (i + 1) % 4 {
				sides++
				sc.RecordVisit(v+sc.dfs.fourDirections[i], v)
			}
			break
		}

		straight := dfs.fourDirections[directionI]
		if w := v + straight; dfs.IsConnected(v, w) && dfs.cc[w] != clusterId {
			v = w
			continue
		} else {
			sc.RecordVisit(w, v)
			sides++
		}

		rightDirectionI := (directionI + 1) % 4
		right := dfs.fourDirections[rightDirectionI]
		if w := v + right; dfs.IsConnected(v, w) && dfs.cc[w] != clusterId {
			v = w
			directionI = rightDirectionI
			continue
		} else {
			sc.RecordVisit(w, v)
			sides++
		}

		// Should return
		directionI = (directionI + 2) % 4
		if w := v + dfs.fourDirections[directionI]; dfs.IsConnected(v, w) && dfs.cc[w] != clusterId {
			v = w
		} else {
			for i := directionI; i != startDirection; i = (i + 1) % 4 {
				sides++
				sc.RecordVisit(v+sc.dfs.fourDirections[directionI], v)
			}
			break
		}
	}

	return sides
}

func (sc *SidesCalc) RecordVisit(v, w int) {
	if !sc.dfs.IsConnected(v, w) {
		w = sc.dfs.n
	}

	sc.visited[v][w] = true
}

func (sc *SidesCalc) FindFirstVertex(clusterId int, start int) int {
	v := -1

	for i := start; i < sc.dfs.n; i++ {
		if sc.dfs.cc[i] == clusterId {
			v = i
			break
		}
	}

	return v
}
