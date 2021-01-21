package main

func canJump(nums []int) bool {
	right, edge := 0, len(nums)-1
	for i := range nums {
		if right < i {
			return false
		}
		t := i + nums[i]
		if t >= edge {
			return true
		}
		if t > right {
			right = t
		}
	}
	return false
}
