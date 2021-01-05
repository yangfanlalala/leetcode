package main

import "sort"

//给定数组 nums = [-1, 0, 1, 2, -1, -4]，
//
//满足要求的三元组集合为：
//[
//[-1, 0, 1],
//[-1, -1, 2]
//]

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0, 1)
Outter:
	for i := 0; i < len(nums); i++ {
		x, y := i+1, len(nums)-1
		for x < y {
			if nums[i]+nums[x]+nums[y] < 0 {
				x++
			} else if nums[i]+nums[x]+nums[y] > 0 {
				y--
			} else {
				result = append(result, []int{nums[i], nums[x], nums[y]})
				continue Outter
			}
		}
	}
	return result
}
