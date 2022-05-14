package reverse

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var prev *ListNode
	cur := head
	for cur != nil {
		if prev == nil { // this is tail
			prev = cur
			cur = cur.Next
			prev.Next = nil
		} else {
			temp := prev
			prev = cur
			cur = cur.Next
			if cur == nil {
				head = prev
			}
			prev.Next = temp
		}
	}
	return head
}
