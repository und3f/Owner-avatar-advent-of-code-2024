package twoD

/*
 *   ^
 * -2|
 * -1|
 *--0+--->
 *  1| 23
 *  2|
 */

var (
	DirectionNorth = []int{-1, 0}
	DirectionSouth = []int{+1, 0}

	DirectionEast = []int{0, +1}
	DirectionWest = []int{0, -1}

	DirectionStay = []int{0, 0}

	DirectionNE = []int{-1, +1}
	DirectionNW = []int{-1, -1}
	DirectionSE = []int{+1, +1}
	DirectionSW = []int{+1, -1}

	CardinalDirections = map[string][]int{
		"n": DirectionNorth,
		"s": DirectionSouth,
		"e": DirectionEast,
		"w": DirectionWest,
	}

	AsciiDirections = map[rune][]int{
		'>': DirectionEast,
		'<': DirectionWest,
		'^': DirectionNorth,
		'v': DirectionSouth,
	}
)

func RotateClockwise(in []int) []int {
	return []int{in[1], -in[0]}
}

func RotateClockwise45(in []int) []int {
	return []int{in[1] + in[0], -in[0] + in[1]}
}

func RotateClockwise135(in []int) []int {
	return []int{in[1] - in[0], -in[0] - in[1]}
}

func RotateCounterclockwise(in []int) []int {
	return []int{-in[1], in[0]}
}

func Reverse(in []int) []int {
	return []int{-in[0], -in[1]}
}
