package _链表

/**
给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。

示例 1：

输入：head = [1,2,3,4]
输出：[2,1,4,3]
示例 2：
输入：head = []
输出：[]
示例 3：
输入：head = [1]
输出：[1]
leetcode: https://leetcode.cn/problems/swap-nodes-in-pairs/description/
提示：

链表中节点的数目在范围 [0, 100] 内
0 <= Node.val <= 100
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//自己写的
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	pre, tmphead, cur := head, head, head.Next
	for {
		tempVal := cur.Val
		cur.Val = pre.Val
		pre.Val = tempVal
		if cur.Next == nil || cur.Next.Next == nil {
			break
		}
		pre = cur.Next
		cur = cur.Next.Next
	}

	return tmphead
}

// 题解 ： 通过交换指针去实现 相邻值翻转
func swapPairs1(head *ListNode) *ListNode {
	viturlHead := &ListNode{
		Next: head,
	}
	vh := viturlHead
	for head != nil && head.Next != nil {
		temp := head.Next
		head.Next = temp.Next
		temp.Next = head
		vh.Next = temp

		//移动 : 注意变换后的移动问题
		vh = head
		head = head.Next
	}

	return viturlHead.Next
}
