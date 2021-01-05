package main

func convert(s string, numRows int) string {
	if numRows == 0 {
		return ""
	}
	if numRows == 1 {
		return s
	}
	result := make([]byte, 0, len(s))
	leng := 2 * (numRows - 1)
	x, y := len(s)/leng, len(s)%leng
	if y > 0 {
		x += 1
	}
	for i := 0; i < numRows; i++ {
		for j := 0; j < x; j++ {
			if j*leng+i < len(s) {
				result = append(result, s[j*leng+i])
			}
			if (i != 0 && i != numRows-1) && (j+1)*leng-i < len(s) {
				result = append(result, s[(j+1)*leng-i])
			}
		}
	}
	return string(result)
}
