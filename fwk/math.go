package fwk

import (
	"math"

	"golang.org/x/exp/constraints"
)

func Abs[V constraints.Signed](a V) V {
	if a < 0 {
		return -a
	}
	return a
}

func CountDigits[V constraints.Integer](num V) V {
	return V(1 + math.Floor(math.Log10(math.Max(2, math.Abs(float64(num))))))
}

func LCM[V constraints.Integer](nums ...V) V {
	res := nums[0]
	for _, num := range nums[1:] {
		res = lcm(res, num)
	}

	return res
}

func lcm[V constraints.Integer](a, b V) V {
	return a * b / gcd(a, b)
}

func GCD[V constraints.Integer](nums ...V) V {
	res := nums[0]
	for _, num := range nums[1:] {
		res = gcd(res, num)
	}

	return res
}

func gcd[V constraints.Integer](a, b V) V {
	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}
