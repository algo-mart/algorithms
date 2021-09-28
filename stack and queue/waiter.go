package main

func primeGenerator(stack []int32) {
	count := 0
	for i := 2; count < len(stack); i++ {
		prime := true
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				prime = false
				break
			}
		}
		if prime {
			stack[count] = int32(i)
			count++
		}
	}
}

func waiter(number []int32, q int32) []int32 {
	stack := make([]int32, q)
	primeGenerator(stack)
	idx := 0
	result := make([]int32, len(number))
	for _, prime := range stack {
		for i, num := range number {
			if num%prime == 0 {
				result[idx] = num
				number[i] = num * -1
				idx++
			}
		}
		newStack := make([]int32, 0)
		for i := len(number) - 1; i >= 0; i-- {
			if number[i] > 0 {
				newStack = append(newStack, number[i])
			}
		}
		number = newStack
	}

	for i := len(number) - 1; i >= 0; i-- {
		result[idx] = number[i]
		idx++
	}
	return result
}

func main() {

}
