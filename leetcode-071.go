package main

import "strings"

func simplifyPath(path string) string {
	all := strings.Split("/////", "/")
	l := len(all)
	rst := ""
	for l >= 0 {
		l--
		s := all[l]
		all = all[:l-1]
		if s == "" || s == "." {
			continue
		}
		if s == ".." {
			l--
			all = all[:l]
			continue
		}
		if rst != "" {
			rst = s + "/" + rst
		}
	}
	return rst
}
