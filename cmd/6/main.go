package main

import (
	"reflect"

	"github.com/und3f/aoc/2024/fwk"
	"github.com/und3f/aoc/2024/fwk/twoD"
)

type Map struct {
	board [][]rune

	cursor    []int
	direction []int
}

func main() {
	input := ReadInput()

	part1(input)
	part2(input)
}

func part1(amap Map) {
	visited := EvaluateBoard(amap)

	fwk.Solution(1, len(visited))
}

func part2(amap Map) {
	var sum int64 = 0

	possibleObstacles := EvaluateBoard(amap)
	for _, obstacle := range possibleObstacles {
		if replacedValue := amap.board[obstacle[0]][obstacle[1]]; !reflect.DeepEqual(obstacle, amap.cursor) && replacedValue != '#' {
			amap.board[obstacle[0]][obstacle[1]] = '#'

			if res := EvaluateBoard(amap); res == nil {
				sum++
			}

			amap.board[obstacle[0]][obstacle[1]] = replacedValue
		}
	}

	fwk.Solution(2, sum)
}

func EvaluateBoard(amap Map) [][]int {
	visited := make([][][]int, len(amap.board))

	for i := range visited {
		visited[i] = make([][]int, len(amap.board[i]))
	}
	stepsFuse := len(amap.board) * len(amap.board)

	cursor := amap.cursor
	direction := amap.direction

	for stepsFuse >= 0 {
		stepsFuse--

		visited[cursor[0]][cursor[1]] = direction
		newCursor := fwk.AddVec(cursor, direction)

		if newCursor[0] < 0 || newCursor[0] >= len(amap.board) ||
			newCursor[1] < 0 || newCursor[1] >= len(amap.board[newCursor[0]]) {
			break
		}

		if amap.board[newCursor[0]][newCursor[1]] == '#' {
			direction = twoD.RotateClockwise(direction)
			continue
		}

		cursor = newCursor

		if oldDirection := visited[cursor[0]][cursor[1]]; oldDirection != nil {
			if reflect.DeepEqual(oldDirection, direction) {
				// fmt.Println("In a loophole")
				return nil
			}
		}
	}

	var possibleObstacles [][]int
	for i := 0; i < len(visited); i++ {
		for j := 0; j < len(visited[i]); j++ {
			if visited[i][j] != nil {
				possibleObstacles = append(possibleObstacles, []int{i, j})
			}
		}
	}

	if stepsFuse < 0 {
		return nil
	}
	return possibleObstacles
}

func ReadInput() Map {
	board := fwk.ReadInputRunesLines()
	cursor := findCursor(board)

	return Map{
		board:     board,
		cursor:    cursor,
		direction: twoD.DirectionNorth,
	}
}

func findCursor(board [][]rune) []int {
	for i, row := range board {
		for j, v := range row {
			if v == '^' {
				return []int{i, j}
			}
		}
	}

	panic("Not found.")
}
