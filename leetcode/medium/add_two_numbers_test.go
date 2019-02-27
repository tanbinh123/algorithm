package medium

import (
	"fmt"
	"testing"
)

/**
2. 两数相加
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	last := 0
	var lastNode *ListNode
	var firstNode *ListNode
	for {
		i, j := 0, 0
		if l1 == nil && l2 == nil && last == 0 {
			break
		}
		if l1 != nil {
			i = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			j = l2.Val
			l2 = l2.Next
		}
		sum := last + i + j
		last = sum / 10
		if lastNode == nil {
			firstNode = &ListNode{
				Val: sum % 10,
			}
			lastNode = firstNode
		} else {
			lastNode.Next = &ListNode{
				Val: sum % 10,
			}
			lastNode = lastNode.Next
		}
	}
	return firstNode
}

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	node := AddTwoNumbers(l1, l2)
	for {
		if node == nil {
			break
		}
		fmt.Println(node.Val)
		node = node.Next
	}
}
