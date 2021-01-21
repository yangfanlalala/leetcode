package main

func lengthOfLastWord(s string) int {
	leng, cnt := len(s), 0
	for i := leng - 1; i >= 0; i-- {
		if s[i] == ' ' && cnt > 0 {
			return cnt
		} else if s[i] != ' ' {
			cnt++
		}
	}
	return cnt
}
