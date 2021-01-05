package main

func largeGroupPositions(s string) [][]int {
	if s == "" {
		return nil
	}
	c := 0
	s = s + "#"
	result := make([][]int, 0, 1)
	for i := 1; i < len(s); i++ {
		if s[c] == s[i] {
			continue
		}
		if i-c >= 3 {
			result = append(result, []int{c, i - 1})
		}
		c = i
	}
	return result
}
