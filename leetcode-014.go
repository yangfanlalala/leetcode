package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	result := make([]byte, 0, len(strs))
	str := strs[0]
	strs = strs[1:]
Out:
	for i := 0; i < len(str); i++ {
		for j := 0; j < len(strs); j++ {
			if i >= len(strs[j]) {
				break Out
			}
			if str[i] != strs[j][i] {
				break Out
			}
		}
		result = append(result, str[i])
	}
	return string(result)
}
