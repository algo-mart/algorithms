package main
/*
https://www.codewars.com/kata/529e2e1f16cb0fcccb000a6b/
 */
func SplitInteger(num int, parts int) []int {
	values:= num/parts
	count := 0
	remainder:=num % parts
	var splittedValues []int
	if num % parts ==0 {
		for count < parts {
			splittedValues= append(splittedValues, values)
			count++
		}
		return splittedValues
	} else {
		for count < parts - remainder {
			splittedValues= append(splittedValues, values)
			count++
		}
		coun:=0
		for coun <  remainder {
			splittedValues = append(splittedValues, values+1)
			coun++
		}
		return splittedValues
	}

}
