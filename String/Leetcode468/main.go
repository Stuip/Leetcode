package main

import (
	"fmt"
	"strconv"
	"strings"
)

func validIPAddress(queryIP string) string {
	isIPv4, isIPv6 := strings.Contains(queryIP, "."), strings.Contains(queryIP, ":")
	if isIPv4 && ipv4Check(queryIP) {
		return "IPv4"
	}
	if isIPv6 && ipv6Check(queryIP) {
		return "IPv6"
	}
	return "Neither"
}

func ipv4Check(ip string) bool {
	// ipv4 分割的长度必须为4
	if items := strings.Split(ip, "."); len(items) == 4 {
		for _, item := range items {
			// 每一个分割字符个数只能是1-3，并且不能有前导零
			if len(item) > 3 || len(item) == 0 || (len(item) > 1 && item[0] == '0') {
				return false
			}
			if v, err := strconv.Atoi(item); err != nil || v > 255 {
				return false
			}
		}
		return true
	} else {
		return false
	}
}

func ipv6Check(ip string) bool {
	// ipv4 分割的长度必须为8
	if items := strings.Split(ip, ":"); len(items) == 8 {
		for _, item := range items {
			// 分割的每个字符串长度在1到4之间
			if len(item) > 4 || len(item) == 0 {
				return false
			}
			// 每个字符都需要在0-9或a-f或A-F中
			if _, err := strconv.ParseUint(item, 16, 64); err != nil {
				return false
			}
		}
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(validIPAddress("172.16.254.1"))
}
