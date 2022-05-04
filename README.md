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

這樣時間複雜度是O(n)

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
     if idx == aLen - 1 {
         head = arr[idx]
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