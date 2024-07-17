package _链表

//Definition for singly-linked list.
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

// 自己的解法
//
//leetcode:https://leetcode.cn/problems/reverse-linked-list/
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var pre *ListNode = nil
	cur := head

	for {
		tem := cur.Next
		cur.Next = pre
		pre = cur
		if tem == nil {
			break
		}
		cur = tem

	}
	return cur
} //return pre 就不用 if 哦安短， 然后 for cur ！= nil  就可以了，牛的

// 题解
// 双指针
func reverseList1(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
