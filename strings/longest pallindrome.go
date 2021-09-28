package main

import "math"

func longestPalindrome(s string) string {
	start, end := 0, 0
	for idx := range s {
		length := int(math.Max(expandRadius(s, idx, idx), expandRadius(s, idx, idx+1)))
		if length > end-start {
			start = idx - (length-1)/2
			end = idx + length/2
		}
	}

	return s[start : end+1]
}
func expandRadius(s string, L int, R int) float64 {
	for L >= 0 && R < len(s) && s[L] == s[R] {
		L--
		R++
	}
	temp := R - L - 1
	return float64(temp)
}

func main() {

}
