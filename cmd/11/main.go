package main

import (
	"fmt"
	"math"
	"runtime"
	"strconv"
	"strings"
	"sync"

	"github.com/und3f/aoc/2024/fwk"
)

type CachableBlinker struct {
	cache map[uint64][]uint64
	step  uint64
}

const (
	PART1_STEPS = 25

	PART2_STEPS = 75
	CACHE_STEP  = 25
)

func main() {
	input := ParseInput()

	part1(input)
	part2(input)
}

func ParseInput() []uint64 {
	var values []uint64

	line := fwk.ReadInputLines()[0]
	for _, str := range strings.Split(line, " ") {
		val, _ := strconv.ParseUint(str, 10, 64)
		values = append(values, val)
	}

	return values
}

func part1(stones []uint64) {
	for i := 0; i < PART1_STEPS; i++ {
		stones = Blink(stones)
	}

	fwk.Solution(1, len(stones))
}

func part2(stones []uint64) {
	threads := runtime.NumCPU()
	stonesOut := make(chan uint64, len(stones))
	stonesIn := make(chan uint64)

	var wg sync.WaitGroup
	for i := 0; i < threads; i++ {
		go func(id int) {
			wg.Add(1)

			defer wg.Done()

			blinker := NewCachableBlinker(CACHE_STEP)

			for stone := range stonesIn {
				stonesOut <- blinker.BlinkNCount(stone, PART2_STEPS)
			}
		}(i)
	}

	for _, stone := range stones {
		stonesIn <- stone
	}

	close(stonesIn)
	wg.Wait()
	close(stonesOut)

	var sum uint64
	for count := range stonesOut {
		sum += count
	}

	fwk.Solution(2, sum)
}

func NewCachableBlinker(step uint64) *CachableBlinker {
	return &CachableBlinker{
		cache: make(map[uint64][]uint64),
		step:  step,
	}
}

func (b *CachableBlinker) BlinkNCount(stone uint64, steps uint64) uint64 {
	if steps%b.step != 0 {
		panic(fmt.Sprintf("Cache step not aligned for %d steps.\n", steps))
	}

	stones, exists := b.cache[stone]

	if !exists {
		stones = BlinkNSingleStone(stone, b.step)

		b.cache[stone] = stones
	}

	stepsToDo := steps - b.step
	if stepsToDo == 0 {
		return uint64(len(stones))
	}

	var sum uint64
	for _, stone := range stones {
		sum += b.BlinkNCount(stone, stepsToDo)
	}
	return sum
}

func BlinkNSingleStone(stone uint64, times uint64) []uint64 {
	return BlinkN([]uint64{stone}, times)
}

func BlinkN(stones []uint64, times uint64) []uint64 {
	for i := uint64(0); i < times; i++ {
		stones = Blink(stones)
	}
	return stones
}

func Blink(_stones []uint64) []uint64 {
	var stones []uint64

	for _, stone := range _stones {
		stones = append(stones, BlinkSingleStone(stone)...)
	}

	return stones
}

func BlinkSingleStone(stone uint64) []uint64 {

	if stone == 0 {
		return []uint64{1}
	} else if digits := fwk.CountDigits(stone); digits%2 == 0 {
		split := uint64(math.Pow(10., float64(digits)/2.))
		return []uint64{stone / split, stone % split}
	} else {
		return []uint64{stone * 2024}
	}

}
