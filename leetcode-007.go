package main

import "math"

func reverse(x int) int {
	lt0 := false
	if x < 0 {
		lt0 = true
		x *= -1
	}
	result := 0
	for x > 0 {
		result = result*10 + x%10
		x /= 10
	}
	if result > math.MaxInt32 {
		return 0
	}
	if lt0 {
		result *= -1
	}
	return result
}
