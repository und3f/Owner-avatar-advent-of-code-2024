package main

import (
	"strings"

	"github.com/und3f/aoc/2024/fwk"
	"github.com/und3f/aoc/2024/fwk/twoD"
)

type Warehouse struct {
	Board     [][]rune
	Movements []rune
	Position  []int
}

type WideWarehouse Warehouse

func main() {
	warehouse, wideWarehouse := ParseInput()

	part1(warehouse)
	part2(wideWarehouse)
}

func part1(warehouse Warehouse) {

	for _, direction := range warehouse.Movements {
		warehouse.Move(direction)
	}

	fwk.Solution(1, warehouse.Score())
}

func part2(warehouse WideWarehouse) {
	for _, direction := range warehouse.Movements {
		// fmt.Printf("Move %c:\n", direction)
		// fmt.Println(fwk.StringifyRunesLines(warehouse.Board))
		warehouse.Move(direction)
	}
	fwk.Solution(2, warehouse.Score())
}

func (w *Warehouse) Move(direction rune) {
	vec := twoD.AsciiDirections[direction]

	movePos := fwk.AddVec(w.Position, vec)

	obj := w.Board[movePos[0]][movePos[1]]
	if obj == '#' {
		return
	}
	if obj == '.' {
		w.Replace(movePos)
		return
	}

	if obj != 'O' {
		panic("Not expected")
	}

	robotMovePos := movePos
	for ; w.Board[movePos[0]][movePos[1]] == 'O'; movePos = fwk.AddVec(movePos, vec) {
	}

	obj = w.Board[movePos[0]][movePos[1]]
	if obj == '#' {
		return
	}
	if obj == '.' {
		w.Replace(robotMovePos)
		w.Board[movePos[0]][movePos[1]] = 'O'
	}
}

func (w *Warehouse) Replace(newpos []int) {
	w.Board[w.Position[0]][w.Position[1]] = '.'
	w.Board[newpos[0]][newpos[1]] = '@'
	w.Position = newpos
}

func (w *Warehouse) Score() uint64 {
	return ScoreWarehouse(w.Board, 'O')
}

func (w *WideWarehouse) Score() uint64 {
	return ScoreWarehouse(w.Board, '[')
}

func (w *WideWarehouse) Move(direction rune) {
	vec := twoD.AsciiDirections[direction]

	newPos := fwk.AddVec(w.Position, vec)
	if w.canPush(newPos, vec) {
		w.push(newPos, vec)
		w.relocate(newPos)
	}
}

func (w *WideWarehouse) relocate(pos []int) {
	w.Board[pos[0]][pos[1]] = '@'
	w.Board[w.Position[0]][w.Position[1]] = '.'
	w.Position = pos
}

func (w *WideWarehouse) canPush(p []int, vec []int) bool {
	v := w.Board[p[0]][p[1]]

	switch v {
	case '.':
		return true
	case '#':
		return false
	case '[', ']':
		if vec[0] == 0 {
			// Horizontal push
			return w.canPush(fwk.AddVec(p, vec), vec)
		} else {
			// Vertical push
			tl := p
			tr := p
			switch v {
			case '[':
				tr = fwk.AddVec(tl, []int{0, 1})
			case ']':
				tl = fwk.AddVec(tl, []int{0, -1})
			}

			return w.canPush(fwk.AddVec(tl, vec), vec) &&
				w.canPush(fwk.AddVec(tr, vec), vec)
		}
	}

	panic("Not expected " + string(v))
}
func (w *WideWarehouse) push(p []int, vec []int) {
	v := w.Board[p[0]][p[1]]

	switch v {
	case '.':
		return
	case '[', ']':
		if vec[0] == 0 {
			// Horizontal push
			moveTo := fwk.AddVec(p, vec)
			w.push(moveTo, vec)
			w.Board[moveTo[0]][moveTo[1]] = w.Board[p[0]][p[1]]
		} else {
			// Vertical push
			tl := p
			tr := p
			switch v {
			case '[':
				tr = fwk.AddVec(tl, []int{0, 1})
			case ']':
				tl = fwk.AddVec(tl, []int{0, -1})
			}

			for _, pos := range [][]int{tl, tr} {
				w.push(fwk.AddVec(pos, vec), vec)

				w.Board[pos[0]+vec[0]][pos[1]] = w.Board[pos[0]][pos[1]]
				w.Board[pos[0]][pos[1]] = '.'
			}
		}
		return
	}

	panic("Should not occur " + string(v))
}

func ParseInput() (Warehouse, WideWarehouse) {
	parts := strings.Split(fwk.ReadInput(fwk.GetDataFilename(1)), "\n\n")

	lines := strings.Split(parts[0], "\n")

	if len(lines[len(lines)-1]) == 0 {
		lines = lines[:len(lines)-1]
	}

	runes := make([][]rune, len(lines))
	for i := range lines {
		runes[i] = []rune(lines[i])
	}

	warehouse := Warehouse{
		Board:     runes,
		Movements: []rune(strings.TrimSpace(strings.Join(strings.Split(parts[1], "\n"), ""))),
		Position:  FindPosition(runes),
	}

	wideRunes := make([][]rune, len(runes))
	for i, row := range runes {
		wideRunes[i] = make([]rune, 2*len(row))
		for j, v := range row {
			var r [2]rune
			switch v {
			case 'O':
				r[0] = '['
				r[1] = ']'
			case '@':
				r[0] = '@'
				r[1] = '.'
			default:
				r[0] = v
				r[1] = v
			}
			wideRunes[i][j*2] = r[0]
			wideRunes[i][j*2+1] = r[1]
		}
	}
	wideWarehouse := WideWarehouse{
		Board:     wideRunes,
		Movements: warehouse.Movements,
		Position:  FindPosition(wideRunes),
	}

	return warehouse, wideWarehouse
}

func ScoreWarehouse(board [][]rune, sym rune) (score uint64) {
	for i, row := range board {
		for j, v := range row {
			if v == sym {
				score += uint64(i*100 + j)
			}
		}
	}

	return
}

func FindPosition(runes [][]rune) []int {
	for i, l := range runes {
		for j, v := range l {
			if v == '@' {
				return []int{i, j}
			}
		}
	}
	panic("Start position not found!")
}
