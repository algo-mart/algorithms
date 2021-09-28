package main

import (
	"math"
	"strings"
)

func intToRoman(num int) string {
	var sb strings.Builder
	for num > 0 {
		placeVal := int(math.Log10(float64(num)))
		multiplier := int(math.Pow10(placeVal))
		digit := (num) / multiplier
		if placeVal == 0 {
			helper(&sb, digit, "I", "V", "IX")
		} else if placeVal == 1 {
			helper(&sb, digit, "X", "L", "XC")
		} else if placeVal == 2 {
			helper(&sb, digit, "C", "D", "CM")
		} else if placeVal == 3 {
			if digit < 4 {
				sb.Write([]byte(strings.Repeat("M", digit)))
			}
		}
		num = (num) % multiplier
	}
	return sb.String()
}
func helper(sb *strings.Builder, digit int, one string, five string, nine string) strings.Builder {
	if digit < 4 {
		sb.Write([]byte(strings.Repeat(one, digit)))
	} else if digit == 4 {
		sb.WriteString(one)
		sb.WriteString(five)
	} else if digit == 9 {
		sb.WriteString(nine)
	} else if digit > 4 {
		sb.Write([]byte(five))
		sb.WriteString(strings.Repeat(one, digit-5))
	}
	return *sb
}

func main() {

}
