package reverse

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseListV1(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	head, _ = ReverseRecursive(head)
	return head
}

func ReverseRecursive(head *ListNode) (*ListNode, *ListNode) {
	if head.Next != nil {
		prevHead, prevTail := ReverseRecursive(head.Next)
		prevTail.Next = head
		head.Next = nil
		return prevHead, head
	} else {
		return head, head
	}
}
