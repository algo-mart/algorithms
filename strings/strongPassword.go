package main

import "fmt"

/*
https://www.hackerrank.com/contests/hourrank-24/challenges/strong-password/problem
 */

func minimumNumber(n int32, password string) int32 {
	// Return the minimum number of characters to make the password strong
	countNum := 0
	countLower := 0
	countUpper :=0
	specialCharacters := 0
	count:=int32(0)
	for _,j := range password {
		if j>='0' && j<='9'{
			countNum++
		} else if j>='a' && j<='z' {
			countLower++
		} else if j>='A' && j<='Z'{
			countUpper++
		} else{
			specialCharacters++
		}
	}
	if countNum==0 {
		count++
	}
	if countLower == 0 {
		count++
	}
	if countUpper == 0 {
		count++
	}
	if specialCharacters==0{
		count++
	}
	balance := 6-(n+count)
	if n+count<6 {
		count = count + balance
	}
	return count

}

func main() {
	fmt.Println('1','9', 'a')
}
