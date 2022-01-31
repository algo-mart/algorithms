package main

import "fmt"

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
//optimized two-number-sum
func TwoNumberSumOptimized(array []int, target int) []int {
	// Write your code here.
	myMap:= make(map[int]bool)

	for i:=0; i<len(array); i++ {
		number:= target-array[i]
		if myMap[number] {
			return [] int {array[i], number}
		} else {
			myMap[array[i]] = true
		}

	}
	return []int{}
}

func main() {
	fmt.Println(TwoNumberSum( [] int {3,5,-4,8,11,1, -1, 6} , 10 ))
	fmt.Println(TwoNumberSumOptimized( [] int {3,5,-4,8,11,1, -1, 6} , 10 ))
}

