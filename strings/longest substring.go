package main

import "math"

func lengthOfLongestSubstring(s string) int {
	myMap := make(map[string]int)
	i := 0
	max := 0
	for j, v := range s {
		if idx, ok := myMap[string(v)]; ok {
			i = int(math.Max(float64(idx), float64(i)))
		}
		myMap[string(v)] = j + 1
		max = int(math.Max(float64(j+1-i), float64(max)))
	}
	return max
}

func main() {

}
