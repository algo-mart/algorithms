package main

import "fmt"

func equalStacks(h1 []int32, h2 []int32, h3 []int32) int32 {
	sum1, sum2, sum3 := int32(0), int32(0), int32(0)
	i := 0
	for i < len(h1) {
		sum1 += h1[i]
		i++
	}
	i = 0
	for i < len(h2) {
		sum2 += h2[i]
		i++
	}
	i = 0
	for i < len(h3) {
		sum3 += h3[i]
		i++
	}
	for len(h1) > 0 && len(h2) > 0 && len(h3) > 0 {
		if sum1 == sum2 && sum2 == sum3 {
			return sum2
		}
		last1 := h1[0]
		last2 := h2[0]
		last3 := h3[0]
		if sum1 < sum2 && sum1 < sum3 {
			h2 = append(h2[1:])

			sum2 -= last2
			h3 = append(h3[1:])
			sum3 -= last3
		} else if sum2 < sum1 && sum2 < sum3 {
			h1 = append(h1[1:])
			sum1 -= last1
			h3 = append(h3[1:])
			sum3 -= last3

		} else if sum3 < sum1 && sum3 < sum2 {
			h1 = append(h1[1:])
			sum1 -= last1
			h2 = append(h2[1:])
			sum2 -= last2
		} else if sum1 < sum2 && sum1 == sum3 {
			h2 = append(h2[1:])
			sum2 -= last2
		} else if sum2 < sum1 && sum2 == sum3 {
			h1 = append(h1[1:])
			sum1 -= last1
		} else if sum2 < sum3 && sum2 == sum1 {
			h3 = append(h3[1:])
			sum3 -= last3
		}

	}
	return 0
}


//solution 2
//LIFO
type Stack struct {
	items [] int32
}

//adds an item to the end of a slice
func (s *Stack) Push(n int32) {
	s.items = append(s.items, n)
}

//remove the value at the end of a
//slice and returns removed value
func (s *Stack) Pop() int32 {
	position := len(s.items)-1
	valueAtPosition := s.items[position]
	s.items = s.items[:position]
	return valueAtPosition
}

func (s *Stack) IsEmpty() bool {
	return len(s.items)==0
}

func (s *Stack) Peek() int32 {
	//if s.IsEmpty() {
	//	//return 0, errors.New("empty slice")
	//}
	value := s.items[len(s.items)-1]
	return value
}


/*
n1[] = {3, 2, 1, 1, 1}
n2[] = {4, 3, 2}
n3[] = {1, 1, 4, 1}
 */
func equalStacks1(h1 []int32, h2 []int32, h3 []int32) int32 {
	// Write your code here

	stack1 := Stack{}
	stack2 := Stack{}
	stack3 := Stack{}

	cumulativeHeight1 := int32(0)
	cumulativeHeight2 := int32(0)
	cumulativeHeight3 := int32(0)

	//fills up stack1 stack2 and stack3
	for i := len(h1)-1; i >=0 ; i-- {
		cumulativeHeight1 +=h1[i]
		stack1.Push(cumulativeHeight1)
		fmt.Println(stack1)
	}
	for i := len(h2)-1; i >=0 ; i-- {
		cumulativeHeight2 +=h2[i]
		stack2.Push(cumulativeHeight2)
		fmt.Println(stack2)
	}
	for i := len(h3)-1; i >=0 ; i-- {
		cumulativeHeight3 +=h3[i]
		stack3.Push(cumulativeHeight3)
		fmt.Println(stack3)
	}

	for true {
		//check if any stack is empty, then answer is 0
		if stack1.IsEmpty() || stack2.IsEmpty() || stack3.IsEmpty() {
			return 0
		}
		cumulativeHeight1 = stack1.Peek()
		cumulativeHeight2 = stack2.Peek()
		cumulativeHeight3 = stack3.Peek()

		if cumulativeHeight1 == cumulativeHeight2 && cumulativeHeight2 == cumulativeHeight3 {
			return cumulativeHeight1
		}
		if cumulativeHeight1 >= cumulativeHeight2 && cumulativeHeight1 >= cumulativeHeight3 {
			stack1.Pop()
		} else if cumulativeHeight2 >= cumulativeHeight3 && cumulativeHeight2 >= cumulativeHeight1 {
			stack2.Pop()
		} else if cumulativeHeight3 >= cumulativeHeight2 && cumulativeHeight3 >= cumulativeHeight1 {
			stack3.Pop()
		}
	}
	return stack1.Peek()
}


func main() {
	h1 := [] int32 {3, 2, 1, 1, 1}
	h2 := [] int32 {4, 3, 2}
	h3 := [] int32 {1, 1, 4, 1}
	fmt.Println(equalStacks1(h1,h2,h3))
	fmt.Println(equalStacks(h1,h2,h3))
}
