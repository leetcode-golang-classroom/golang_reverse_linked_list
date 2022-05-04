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
	arr := []*ListNode{}
	cur := head
	for cur != nil {
		arr = append(arr, cur)
		cur = cur.Next
	}
	aLen := len(arr)
	for idx := aLen - 1; idx >= 0; idx-- {
		if idx == 0 {
			arr[idx].Next = nil
		} else {
			arr[idx].Next = arr[idx-1]
		}
		if idx == aLen-1 {
			head = arr[idx]
		}
	}
	return head
}
