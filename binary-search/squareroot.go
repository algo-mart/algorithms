package main

import (
	"fmt"
	"math"
)
/*
https://leetcode.com/problems/sqrtx
 */
func mySqrt(x int) int {
	left := 1
	right := x
	if x < 2 {
		return x
	}
	//1 2 3 4 5 6 7 8 9 10
	for left < right {
		mid := left + int(math.Floor(float64((right - left)/2)))
		if mid * mid == x {
			return mid
		} else if mid * mid > x {
			right = mid
		} else if mid * mid < x {
			left = mid + 1
		}
	}

	return left - 1
}


func main() {
	fmt.Println(mySqrt(10))
}
