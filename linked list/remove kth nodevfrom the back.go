package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	current := &ListNode{}
	current.Next = head
	first := current
	second := current

	for i := 0; i < n+1; i++ {
		first = first.Next
	}
	for first != nil {
		first = first.Next
		second = second.Next
	}
	second.Next = second.Next.Next
	return current.Next
}

func main() {

}
