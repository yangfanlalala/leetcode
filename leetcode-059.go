package main

func generateMatrix(n int) [][]int {
	result := make([][]int, 0, n)
	for i := 0; i < n; i++ {
		result = append(result, make([]int, n, n))
	}
	top, bottom, left, right := 0, n-1, 0, n-1
	c, m := 1, n*n
	for c <= m {
		for i := left; i <= right; i++ {
			result[top][i] = c
			c++
		}
		top++
		for i := top; i <= bottom; i++ {
			result[i][right] = c
			c++
		}
		right--
		for i := right; i >= left; i-- {
			result[bottom][i] = c
			c++
		}
		bottom--
		for i := bottom; i >= top; i-- {
			result[i][left] = c
			c++
		}
		left++
	}
	return result
}
