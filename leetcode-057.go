package main

func insert(intervals [][]int, newInterval []int) [][]int {
	result := make([][]int, 0, 1)
	left, right, merge := newInterval[0], newInterval[1], false
	for _, i := range intervals {
		if i[0] > right {
			if !merge {
				result = append(result, []int{left, right})
				merge = true
			}
			result = append(result, i)
		} else if i[1] < left {
			result = append(result, i)
		} else {
			left = min(i[0], left)
			right = max(i[1], right)
		}
	}
	if !merge {
		result = append(result, []int{left, right})
	}
	return result
}
