package main

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0, 1)
	for i:=0; i<len(nums)-3; i++ {
		if i>0 && nums[i] == nums[i-1] {
			continue
		}
		for j:=i+1; j<len(nums)-2; j++ {
			if j>i+1 && nums[j] == nums[j-1] {
				continue
			}
			x, y := j+1, len(nums)-1
			for x<y {
				if x>j+1 && nums[x] == nums[x-1] {
					x++
					continue
				}
				sum := nums[i]+nums[j]+nums[x]+nums[y]
				if sum == target {
					result = append(result, []int{nums[i], nums[j], nums[x], nums[y]})
					x++
					y--
				} else if target > sum {
					x++
				} else {
					y--
				}
			}
		}
	}
	return result
}
