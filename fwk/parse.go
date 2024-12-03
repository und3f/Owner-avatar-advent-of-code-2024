package fwk

import (
	"regexp"
	"strconv"

	"golang.org/x/exp/constraints"
)

func ParseInts[V constraints.Integer](line string) []V {
	re := regexp.MustCompile(",\\s*")
	s := re.Split(line, -1)

	ints := make([]V, len(s))
	for i := range s {
		v, err := strconv.Atoi(s[i])
		if err != nil {
			panic(err)
		}

		ints[i] = V(v)
	}

	return ints
}
