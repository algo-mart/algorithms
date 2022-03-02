package main

import (
	"fmt"
	"strings"
)

/*
Input: haystack = "hello", needle = "ll"
Output: 2
 */



func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	val := strings.Contains(haystack, needle)
	if val {
		for i := 0; i < len(haystack); i++ {
			for j := 0; j < len(needle); j++ {
				if needle[j]==haystack[i] {
					return i
				}
			}
		}
	}
	return -1
}
func strStr2(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	num:= strings.Index(haystack, needle)
	if num != -1 {
		return num
	} else {
		return -1
	}
}

/*
"mississippi"
"issip"
 */

func main() {
	fmt.Println(strStr("mississippi","issip"))
}