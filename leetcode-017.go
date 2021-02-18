package main

func letterCombinations(digits string) []string {
	xMap := map[byte][]byte{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'v', 'y', 'z'},
	}
	result := make([]string, 0, 4*len(digits))
	tmp := make([]byte, 0, 4)
	var combine func(i int)
	combine = func(i int) {
		//处理完了最后一个字符
		if i == len(digits) {
			t := make([]byte, len(tmp), len(tmp))
			copy(t, tmp)
			result = append(result, string(t))
			return
		}
		x := xMap[digits[i]]
		for _, b := range x {
			tmp = append(tmp, b)
			combine(i + 1)
			tmp = tmp[:len(tmp)-1]
		}
	}
	combine(0)
	return result
}
