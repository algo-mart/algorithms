package main

// o-n-square
func TwoNumberSum(array []int, target int) []int {
	// Write your code here.
	for i:=0; i<len(array); i++{
		for j:=i+1; j<len(array); j++ {
			if array[i]+array[j]==target {
				return [] int {array[j], array[i]}
			}
		}
	}
	return [] int{}
}


