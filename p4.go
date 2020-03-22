package main

import (
	"math"
	"strconv"
)

func P4(param int) int {
	reverse := func(s string) string {
		rs := []rune(s)
		for i, j := 0, len(rs)-1; i < j; i, j = i+1, j-1 {
			rs[i], rs[j] = rs[j], rs[i]
		}
		return string(rs)
	}

	min := int(math.Pow(10.0, float64(param-1)))
	max := int(math.Pow(10.0, float64(param))) - 1

	answer := 0
	for a := max; a >= min; a-- {
		for b := max; b >= min; b-- {
			ab := strconv.Itoa(a * b)
			if ab == reverse(ab) && answer < a*b {
				answer = a * b
			}
		}
	}
	return answer
}
