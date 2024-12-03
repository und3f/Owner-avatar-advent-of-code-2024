package fwk

import "golang.org/x/exp/constraints"

func checkVectorsMatch[V constraints.Integer](a, b []V) {
	if len(a) != len(b) {
		panic("Vectors not match")
	}
}

func AddVec[V constraints.Integer](a, b []V) []V {
	checkVectorsMatch(a, b)

	r := make([]V, len(a))
	for i := range a {
		r[i] = a[i] + b[i]
	}

	return r
}

func MultVecByConstant[V constraints.Integer](a []V, c V) []V {
	r := make([]V, len(a))
	for i, v := range a {
		r[i] = v * c
	}

	return r
}

func SubVec[V constraints.Signed](a, b []V) []V {
	checkVectorsMatch(a, b)

	r := make([]V, len(a))
	for i := range a {
		r[i] = a[i] - b[i]
	}

	return r
}

func CalManhattan[V constraints.Signed](a, b []V) V {
	var manhattan V = 0

	for _, v := range SubVec(a, b) {
		manhattan += Abs(v)
	}

	return manhattan
}

func AbsVec[V constraints.Signed](a []V) V {
	var sum V = 0

	for i := range a {
		sum += Abs(a[i])
	}

	return sum
}
