package main

import "math"

type MinStack struct {
	items []float64
}

var stack MinStack
var minStack MinStack

//func Constructor() MinStack {
//
//}


func (this *MinStack) Push(val int)  {

	minimumValue := float64(0)
	stack.items = append(stack.items, float64(val))
	if len(minStack.items) != 0 {
		minimumValue = math.Min(float64(val), float64(len(minStack.items)-1))
		minStack.items = append(minStack.items, minimumValue)
	}
}


func (this *MinStack) Pop()  {
	lastItem := len(stack.items)-1
	lastItem2 := len(stack.items)-1
	stack.items = stack.items[:lastItem]
	minStack.items = minStack.items[:lastItem2]
}


func (this *MinStack) Top() int {
	valueAtTop := this.items[len(stack.items)-1]
	return int(valueAtTop)
}


func (this *MinStack) GetMin() int {
	value := minStack.items[len(minStack.items)-1]
	return int(value)
}

