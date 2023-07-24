package main

import (
	"fmt"
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	n := len(s)
	ans, path := []string{}, []string{}
	var dfs func(start int)
	dfs = func(start int) {
		if len(path) == 4 && start == n {
			ans = append(ans, strings.Join(path, "."))
			return
		}
		if len(path) == 4 && start < n {
			return
		}
		for length := 1; length <= 3; length += 1 {
			if start+length-1 >= n {
				return
			}
			if length != 1 && s[start] == '0' {
				return
			}
			str := s[start : start+length]
			if n, _ := strconv.Atoi(str); n > 255 {
				return
			}
			path = append(path, s[start:start+length])
			dfs(start + length)
			path = path[:len(path)-1]
		}

	}
	dfs(0)
	return ans
}

func main() {
	fmt.Println(restoreIpAddresses("25525511135"))
}
