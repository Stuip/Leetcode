package main

import "fmt"

func isValid(s string) bool {
	stack := []byte{}
	n := len(s)
	for i := 0; i < n; i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, s[i])
		}
		if s[i] == ')' {
			if len(stack) != 0 && stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
		if s[i] == '}' {
			if len(stack) != 0 && stack[len(stack)-1] == '{' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
		if s[i] == ']' {
			if len(stack) != 0 && stack[len(stack)-1] == '[' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("(]"))
}
