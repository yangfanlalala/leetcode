package main

import (
	"strconv"
	"strings"
)

//判断一个数字字符串（只包含0-9）是不是有效的

func SplitAllIP(s string) []string {
	if s == "" || len(s) > 12 {
		return nil
	}
	var results = make([]string, 0, 1)
	var tmp = make([]string, 0, 4)
	var px func(string)
	px = func(sx string) {
		if sx == "" {
			if len(tmp) != 4 {
				return
			}
			results = append(results, strings.Join(tmp, "."))
			return
		}
		//传参不为空，且已到达4个分段
		if len(tmp) == 4 {
			return
		}
		for i := 0; i < 3 && i < len(sx); i++ {
			x := sx[:i+1]
			if !VerifyIPSection(x) {
				break
			}
			tmp = append(tmp, x)
			px(sx[i+1:])
			tmp = tmp[:len(tmp)-1]
		}
	}
	px(s)
	return results
}

func VerifyIPSection(s string) bool {
	if s == "" {
		return false
	}
	if len(s) > 1 && strings.HasPrefix(s, "0") {
		return false
	}
	num, _ := strconv.Atoi(s)
	if num > 255 {
		return false
	}
	return true
}
