package main

import "math"

func maxArea(height []int) int {
	leng := len(height)
	x, y, mArea := 0, leng-1, math.MinInt32
	for x < y {
		area := (y - x) * min(height[x], height[y])
		if area > mArea {
			mArea = area
		}
		if height[y] < height[x] {
			y--
		} else {
			x++
		}
	}
	return mArea
}
