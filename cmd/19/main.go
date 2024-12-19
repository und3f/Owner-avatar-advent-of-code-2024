package main

import (
	"fmt"
	"strings"

	"github.com/und3f/aoc/2024/fwk"
)

func main() {
	patterns, towels := ParseInput()
	Part1(patterns, towels)
	Part2(patterns, towels)
}

func Part1(patterns, towels []string) {
	count := 0
	for _, towel := range towels {
		if possible := IsDesignPossibleInfinitePatterns(patterns, towel); possible {
			count++
		}
	}

	fwk.Solution(1, count)
}

func Part2(patterns, towels []string) {
	var count uint64

	cache := make(map[string]uint64)
	dc := &DesignsCounter{patterns, cache}

	for _, towel := range towels {
		count += dc.CountPossibleDesigns(towel)
		fmt.Printf("%s = %d\n", towel, dc.CountPossibleDesigns(towel))
	}
	fwk.Solution(1, count)
}

func IsDesignPossibleInfinitePatterns(patterns []string, towel string) bool {
	if len(towel) == 0 {
		return true
	}

	for _, pat := range patterns {
		if strings.HasPrefix(towel, pat) {
			if IsDesignPossibleInfinitePatterns(patterns, towel[len(pat):]) {
				return true
			}
		}
	}

	return false
}

type DesignsCounter struct {
	patterns []string
	cache    map[string]uint64
}

func (counter *DesignsCounter) CountPossibleDesigns(towel string) uint64 {
	if len(towel) == 0 {
		return 1
	}

	if c, exist := counter.cache[towel]; exist {
		return c
	}

	var count uint64
	for _, pat := range counter.patterns {
		if strings.HasPrefix(towel, pat) {
			pc := counter.CountPossibleDesigns(towel[len(pat):])
			count += pc
		}
	}

	counter.cache[towel] = count
	return count
}

func IsDesignPossible(patterns []string, towel string) bool {
	bfs := &BFSDesignSearch{
		patterns,
		make([]bool, len(patterns)),
	}

	return bfs.SearchPath(towel)
}

type BFSDesignSearch struct {
	patterns []string
	visited  []bool
}

func (bfs *BFSDesignSearch) SearchPath(towel string) bool {
	if len(towel) == 0 {
		return true
	}

	for i, pat := range bfs.patterns {
		if bfs.visited[i] {
			continue
		}

		if !strings.HasPrefix(towel, pat) {
			continue
		}

		bfs.visited[i] = true
		if bfs.SearchPath(towel[len(pat):]) {
			return true
		}

		bfs.visited[i] = false
	}

	return false
}

func ParseInput() ([]string, []string) {
	content := fwk.ReadInput("")
	parts := strings.Split(content, "\n\n")

	patterns := strings.Split(parts[0], ", ")
	towels := strings.Split(strings.TrimSpace(parts[1]), "\n")
	return patterns, towels
}
