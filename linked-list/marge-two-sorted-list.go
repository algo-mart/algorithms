package main

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy
	for l1 != nil || l2 != nil {
		if l1 == nil {
			current.Next = l2
			l2 = l2.Next
			current = current.Next
		} else if l2 == nil {
			current.Next = l1
			l1 = l1.Next
			current = current.Next
		} else {
			if l1.Val < l2.Val {
				current.Next = l1
				current = current.Next
				l1 = l1.Next
			} else {
				current.Next = l2
				current = current.Next
				l2 = l2.Next
			}
		}
	}
	return dummy.Next
}
func main() {

}
