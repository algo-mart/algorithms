package main

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	prev := dummy
	first := dummy.Next

	for first != nil && first.Next != nil {
		third := first.Next.Next
		second := first.Next

		prev.Next = second
		second.Next = first
		first.Next = third

		prev = first
		first = third
	}
	return dummy.Next
}

func main() {

}
