package main

func intToRoman(num int) string {
	xMap := map[int][]byte{
		1:    {'I'},
		4:    {'I', 'V'},
		5:    {'V'},
		9:    {'I', 'X'},
		10:   {'X'},
		40:   {'X', 'L'},
		50:   {'L'},
		90:   {'X', 'C'},
		100:  {'C'},
		400:  {'C', 'D'},
		500:  {'D'},
		900:  {'C', 'M'},
		1000: {'M'},
	}
	numArray := make([]int, 0, 4)
	times := 1
	for num > 0 {
		mod := num % 10
		num = num / 10
		numArray = append(numArray, mod)
		times *= 10
	}
	times /= 10
	result := make([]byte, 0, 3*len(numArray))
	for i := len(numArray) - 1; i >= 0; i-- {
		if numArray[i] == 0 {
			times /= 10
			continue
		}
		x, ok := xMap[numArray[i]*times]
		if ok {
			result = append(result, x...)
		} else {
			y := xMap[times]
			if numArray[i] < 4 {
				for j := 0; j < numArray[i]; j++ {
					result = append(result, y...)
				}
			} else if numArray[i] > 5 {
				result = append(result, xMap[5*times]...)
				for j := 5; j < numArray[i]; j++ {
					result = append(result, y...)
				}
			}
		}
		times /= 10
	}
	return string(result)
}
