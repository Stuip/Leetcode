package main

import "fmt"

func decodeString(s string) string {
	stack_res, stack_multi := []string{}, []int{}
	multi, res := 0, ""
	for _, c := range s {
		if c == '[' {
			stack_multi = append(stack_multi, multi)
			stack_res = append(stack_res, res)
			multi, res = 0, ""
		} else if c == ']' {
			cur_multi, cur_res := stack_multi[len(stack_multi)-1], stack_res[len(stack_res)-1]
			stack_multi = stack_multi[:len(stack_multi)-1]
			stack_res = stack_res[:len(stack_res)-1]
			for i := 0; i < cur_multi; i++ {
				cur_res += res
			}
			res = cur_res
		} else if c >= '0' && c <= '9' {
			multi = multi*10 + int(c-'0')
		} else {
			res += string(c)
		}
	}
	return res
}

func main() {
	fmt.Println(decodeString("3[a]2[bc]"))
}
