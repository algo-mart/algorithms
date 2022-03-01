package main

import (
	"strconv"
	"strings"
)

/*
https://www.codewars.com/kata/52dca71390c32d8fb900002b/

Task
Finish the function numberToOrdinal, which should take a
number and return it as a string with the correct ordinal
indicator suffix (in English). For example, 1 turns into "1st".

For the purposes of this challenge, you may assume that the
function will always be passed a non-negative integer. If the
function is given 0 as an argument, it should return the string "0" without a suffix.

To help you get started, here is an excerpt from Wikipedia's page on Ordinal Indicators:

st is used with numbers ending in 1 (e.g. 1st, pronounced first)
nd is used with numbers ending in 2 (e.g. 92nd, pronounced ninety-second)
rd is used with numbers ending in 3 (e.g. 33rd, pronounced thirty-third)
As an exception to the above rules, all the "teen" numbers ending with 11, 12 or 13 use -th (e.g. 11th, pronounced eleventh, 112th, pronounced one hundred [and] twelfth)
th is used for all other numbers (e.g. 9th, pronounced ninth).
 */

func NumberToOrdinal(number int) string {
	num := strconv.Itoa(number)
	if num>="11" && num <="13"  {
		return num+"th"
	} else if strings.HasSuffix(num, "1")  {
		return num+"st"
	} else if strings.HasSuffix(num, "2"){
		return num+"nd"
	} else if strings.HasSuffix(num, "3") {
		return num+"rd"
	} else if num == "0" {
		return num
	}
	return num+"th"
}
