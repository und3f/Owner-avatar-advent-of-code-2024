package main

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/und3f/aoc/2024/fwk"
)

const (
	WIDTH  = 101
	HEIGHT = 103

	SECONDS_P1 = 100
)

type Robot struct {
	Position []int
	Velocity []int
}

func main() {
	input := ParseInput()

	part1(input)
	part2(input)
}

func part1(robots []Robot) {
	for turn := 0; turn < SECONDS_P1; turn++ {
		MoveRobots(robots)
	}

	mx := int(WIDTH / 2)
	my := int(HEIGHT / 2)

	quadrantCount := [4]uint{}
	for _, r := range robots {
		p := r.Position
		if p[0] == mx || p[1] == my {
			continue
		}

		q := 0
		if p[0] > mx {
			q++
		}
		if p[1] > my {
			q += 2
		}
		quadrantCount[q]++
	}

	res := uint(1)
	for _, c := range quadrantCount {
		res *= c
	}

	fwk.Solution(1, res)
}

func part2(robots []Robot) {
	turn := SECONDS_P1 + 1
	for ; turn <= 100000; turn++ {
		MoveRobots(robots)

		if IsEasterEgg(robots) {
			break
		}
	}

	fwk.Solution(2, turn)
}

func MoveRobots(robots []Robot) {
	for i, r := range robots {
		robots[i].Position = []int{
			(r.Position[0] + r.Velocity[0] + WIDTH) % WIDTH,
			(r.Position[1] + r.Velocity[1] + HEIGHT) % HEIGHT,
		}
	}
}

func IsEasterEgg(robots []Robot) bool {
	grid := fwk.NewCustomInfiniteGrid[rune](' ', func(v rune) string { return string(v) })

	for _, r := range robots {
		grid.SetAt([]int{r.Position[1], r.Position[0]}, 'O')
	}
	img := grid.String()

	// Search for part of ASCII frame
	return strings.Contains(img, "OOOOOOOOOOOOOO")
}

var robotRe = regexp.MustCompile(`^p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)$`)

func ParseInput() []Robot {
	content := fwk.ReadInputLines()
	robots := make([]Robot, len(content))

	for i, line := range content {
		m := robotRe.FindStringSubmatch(line)

		px, _ := strconv.Atoi(m[1])
		py, _ := strconv.Atoi(m[2])

		vx, _ := strconv.Atoi(m[3])
		vy, _ := strconv.Atoi(m[4])

		robots[i] = Robot{[]int{px, py}, []int{vx, vy}}
	}

	return robots
}
