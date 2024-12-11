package fwk

import (
	"bytes"
	"fmt"

	"golang.org/x/exp/constraints"
)

type Grid[V constraints.Integer] interface {
	GetAt(pos []int) V
	SetAt(pos []int, value V)
	CountAll(value V) uint
	String() string
	FindDimensions() (tl [2]int, br [2]int)
}

type InfiniteGrid[V constraints.Integer] struct {
	items        map[int]map[int]V
	defaultValue V
	stringer     func(V) string
}

func NewInfiniteGrid[V constraints.Integer]() Grid[V] {
	return &InfiniteGrid[V]{
		items:        make(map[int]map[int]V),
		defaultValue: '.',
		stringer:     func(v V) string { return fmt.Sprintf("%c", v) },
	}
}

func NewCustomInfiniteGrid[V constraints.Integer](defaultValue V, stringer func(V) string) Grid[V] {
	return &InfiniteGrid[V]{
		items:        make(map[int]map[int]V),
		defaultValue: defaultValue,
		stringer:     stringer,
	}
}

func (grid *InfiniteGrid[V]) GetAt(pos []int) V {
	l, exists := grid.items[pos[0]]
	if !exists {
		return grid.defaultValue
	}

	if v, exists := l[pos[1]]; exists {
		return v
	}
	return grid.defaultValue
}

func (grid *InfiniteGrid[V]) SetAt(pos []int, value V) {

	l, exists := grid.items[pos[0]]
	if !exists {
		l = make(map[int]V)
		grid.items[pos[0]] = l
	}

	if value == grid.defaultValue {
		delete(l, pos[1])

		if len(l) == 0 {
			delete(grid.items, pos[0])
		}
		return
	}

	l[pos[1]] = value
}

func (grid *InfiniteGrid[V]) CountAll(value V) uint {
	var count uint

	tl, br := grid.FindDimensions()
	for y := tl[0]; y <= br[0]; y++ {
		for x := tl[1]; x <= br[1]; x++ {
			if grid.GetAt([]int{y, x}) == value {
				count++
			}
		}
	}

	return count
}

func (grid *InfiniteGrid[V]) String() string {
	var buf bytes.Buffer

	tl, br := grid.FindDimensions()
	for y := tl[0]; y <= br[0]; y++ {
		for x := tl[1]; x <= br[1]; x++ {
			buf.WriteString(grid.stringer(grid.GetAt([]int{y, x})))
		}
		buf.WriteRune('\n')
	}

	return buf.String()
}

func (grid *InfiniteGrid[V]) FindDimensions() (tl [2]int, br [2]int) {
	for y, line := range grid.items {
		tl[0] = min(tl[0], y)
		br[0] = max(br[0], y)

		for x := range line {
			tl[1] = min(tl[1], x)
			br[1] = max(br[1], x)
		}
	}

	return
}
