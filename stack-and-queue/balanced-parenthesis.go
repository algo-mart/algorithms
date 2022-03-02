package main

import (
	"fmt"
	"math"
)

func isBalanced(s string) string {
	// Write your code here
	if s == "" {
		return "YES"
	}

	if len(s)%2 == 1 {
		// fmt.Println("mdjdidi")
		return "NO"
	}

	store := []string{}

	for _, val := range s {
		str := string(val)

		if str == "(" || str == "{" || str == "[" {
			store = append(store, str)
		} else {
			lastString := store[len(store)-1]
			fmt.Println("here",lastString)
			if str == ")" {
				if lastString == "(" {
					store = store[:len(store)-1]
				}
			} else if str == "}" {
				if lastString == "{" {
					store = store[:len(store)-1]
				}
			} else if str == "]" {
				if lastString == "[" {
					store = store[:len(store)-1]
				}
			} else {
				return "NO"
			}
			fmt.Println("here",lastString)
		}
	}

	if len(store) == 0 && len(s)%2 == 0 {
		return "YES"
	}
	return "NO"
}

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
	fmt.Println(isBalanced("((()))"))
	//fmt.Println(isBalanced("(())("))
	fmt.Println(mySqrt(10))
}
