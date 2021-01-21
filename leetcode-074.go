package main

import "math"

func searchMatrix(matrix [][]int, target int) bool {
	n := len(matrix)
	if n == 0 {
		return false
	}
	m := len(matrix[0])
	if m == 0 {
		return false
	}
	x, y := 0, n*m-1
	for x <= y {
		mid := int(math.Ceil(float64(x+y) / 2))
		i, j := mid/m, mid%m
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] < target {
			x += 1
		} else {
			y -= 1
		}
	}
	return false
}
