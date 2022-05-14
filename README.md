# golang_reverse_linked_list

Given the `head`
 of a singly linked list, reverse the list, and return *the reversed list*

## Examples

**Example 1:**

![https://assets.leetcode.com/uploads/2021/02/19/rev1ex1.jpg](https://assets.leetcode.com/uploads/2021/02/19/rev1ex1.jpg)

```
Input: head = [1,2,3,4,5]
Output: [5,4,3,2,1]

```

**Example 2:**

```
Input: head = [1,2]
Output: [2,1]

```

**Example 3:**

```
Input: head = []
Output: []

```

**Constraints:**

- The number of nodes in the list is the range `[0, 5000]`.
- `5000 <= Node.val <= 5000`

## 解析

題目給定一個單向的鏈結串列 。要實做出一個演算法把鏈結串列反轉

這題的關鍵點在於如何記住已經走過的結點做轉向

比較簡單的作法就是用一個指標陣列把走過的結點都紀錄下來

然後在把陣列倒著串起來

這樣時間複雜度是O(n) 而空間複雜度也會是 O(n)

而透過另一個方式

每次把走訪到的結點都放到最前面，多用一個指標來記住上一個結點

這樣走訪過一次方好可以反向所有的結點

這樣時間複雜度是O(n) 會降低到 O(1)

![](https://i.imgur.com/UefJteX.png)

如果遞迴的作法可以觀察下圖

![](https://i.imgur.com/T55TFoo.png)

## 程式碼

```go
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
```

## 困難點

1. 要理解鏈結串列的特性
2. 要設計一個資料結構可以記住走過的結點

## Solve Point

- [x]  Understand what the problem would like to solve
- [x]  Analysis Complexity