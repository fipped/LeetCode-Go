package leetcode

import (
	"github.com/halfrost/LeetCode-Go/structures"
)

// ListNode define
type ListNode = structures.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 1. 把结果存到 l1。
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    head := l1
    for ; l1 != nil && l2 != nil; l1, l2 = l1.Next, l2.Next {
        l1.Val += l2.Val
        addToNext(l1)
        if l1.Next == nil {
            l1.Next, l2.Next = l2.Next, l1.Next
        }
    }
    for ; l1 != nil && l1.Val >= 10; l1 = l1.Next {
        addToNext(l1)
    }
    return head
}

// 进行进位。
func addToNext(l1 *ListNode) {
    ten := l1.Val / 10
    l1.Val %= 10
    if l1.Next != nil {
        l1.Next.Val += ten
    } else if ten > 0 {
        l1.Next = &ListNode{Val: ten}
    }
}

// 2. 新开变量存结果，模拟加法。
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Val: 0}
	n1, n2, carry, current := 0, 0, 0, head
	for l1 != nil || l2 != nil || carry != 0 {
		if l1 == nil {
			n1 = 0
		} else {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 == nil {
			n2 = 0
		} else {
			n2 = l2.Val
			l2 = l2.Next
		}
		current.Next = &ListNode{Val: (n1 + n2 + carry) % 10}
		current = current.Next
		carry = (n1 + n2 + carry) / 10
	}
	return head.Next
}
