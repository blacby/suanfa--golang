package _链表

//给你一个链表的头节点 head 和一个整数 val ，请你删除链表中所有满足 Node.val == val 的节点，并返回 新的头节点
//示例 1：
//
//
//输入：head = [1,2,6,3,4,5,6], val = 6
//输出：[1,2,3,4,5]
//示例 2：
//
//输入：head = [], val = 1
//输出：[]
//示例 3：
//
//输入：head = [7,7,7,7], val = 7
//输出：[]
//https://leetcode.cn/problems/remove-linked-list-elements/description/

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 自己思考的方式
//
//	func removeElements(head *ListNode, val int) *ListNode {
//		if head == nil {
//			return nil
//		}
//		preHead := &ListNode{head.Val, head}
//		for head != nil && head.Val == val {
//			if head.Val == val {
//				preHead.Next = head.Next
//				head = head.Next
//			}
//		}
//		for head != nil && head.Next != nil {
//			if head.Next.Val == val {
//				head.Next = head.Next.Next
//			} else {
//				head = head.Next
//			}
//		}
//
//		return preHead.Next
//	}
//
// 参考题解的  虚拟头节点的方式
func removeElements(head *ListNode, val int) *ListNode {
	viturlHead := &ListNode{0, head}
	cur := viturlHead
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	return viturlHead.Next
}
