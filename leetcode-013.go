package main

//I             1
//V             5
//X             10
//L             50
//C             100
//D             500
//M             1000

func romanToInt(s string) int {
	xMap := map[byte]int{
		'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000,
	}
	lastNum, result := 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if xMap[s[i]] < lastNum {
			result -= xMap[s[i]]
		} else {
			result += xMap[s[i]]
		}
		lastNum = xMap[s[i]]
	}
	return result
}
