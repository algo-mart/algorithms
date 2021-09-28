package main

func myAtoi(s string) int {
	sign := 1
	count := 0
	number := 0
	for i, x := range s {
		if (x < '0' || x > '9') && x != '-' && x != '+' && x != ' ' {
			return 0
		}
		if x == ' ' {
			continue
		}
		if x == '+' {
			count++
		}
		if x == '-' {
			sign = -1
			count++
		}

		if count < 2 && ((x == '+' || x == '-') || (x >= '0' && x <= '9')) {
			if x >= '0' && x <= '9' {
				number = int(x) - '0'
			}
			i++
			for i < len(s) && s[i] >= '0' && s[i] <= '9' {
				number = number*10 + int(s[i]) - '0'
				i++
				if number >= 2147483648 {
					return (2147483648 * sign) - (sign+1)/2
				}
			}
			return number * sign
		}
	}
	return 0
}

func main() {

}
