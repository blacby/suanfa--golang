package _链表

//leetcode链接：https://leetcode.cn/problems/design-linked-list/description/
//你可以选择使用单链表或者双链表，设计并实现自己的链表。
//
//单链表中的节点应该具备两个属性：val 和 next 。val 是当前节点的值，next 是指向下一个节点的指针/引用。
//
//如果是双向链表，则还需要属性 prev 以指示链表中的上一个节点。假设链表中的所有节点下标从 0 开始。
//
//实现 MyLinkedList 类：
//
//MyLinkedList() 初始化 MyLinkedList 对象。
//int get(int index) 获取链表中下标为 index 的节点的值。如果下标无效，则返回 -1 。
//void addAtHead(int val) 将一个值为 val 的节点插入到链表中第一个元素之前。在插入完成后，新节点会成为链表的第一个节点。
//void addAtTail(int val) 将一个值为 val 的节点追加到链表中作为链表的最后一个元素。
//void addAtIndex(int index, int val) 将一个值为 val 的节点插入到链表中下标为 index 的节点之前。如果 index 等于链表的长度，那么该节点会被追加到链表的末尾。如果 index 比长度更大，该节点将 不会插入 到链表中。
//void deleteAtIndex(int index) 如果下标有效，则删除链表中下标为 index 的节点。
//
//
//示例：
//
//输入
//["MyLinkedList", "addAtHead", "addAtTail", "addAtIndex", "get", "deleteAtIndex", "get"]
//[[], [1], [3], [1, 2], [1], [1], [1]]
//输出
//[null, null, null, null, 2, null, 3]
//
//解释
//MyLinkedList myLinkedList = new MyLinkedList();
//myLinkedList.addAtHead(1);
//myLinkedList.addAtTail(3);
//myLinkedList.addAtIndex(1, 2);    // 链表变为 1->2->3
//myLinkedList.get(1);              // 返回 2
//myLinkedList.deleteAtIndex(1);    // 现在，链表变为 1->3
//myLinkedList.get(1);              // 返回 3
//
//
//提示：
//
//0 <= index, val <= 1000
//请不要使用内置的 LinkedList 库。
//调用 get、addAtHead、addAtTail、addAtIndex 和 deleteAtIndex 的次数不超过 2000 。

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

//看题解后改进，不用创造一个尾巴节点， 并且每次都把index遍历然后退出条件是index <0 &&head next ！= 虚头 牛的

// 自己做的
type MyLinkedList struct {
	Val   int
	Next  *MyLinkedList
	prev  *MyLinkedList
	lenth int
}

func Constructor() MyLinkedList {
	tail := &MyLinkedList{0, nil, nil, 0}
	preHead := &MyLinkedList{0, tail, tail, 0}
	tail.Next = preHead
	tail.prev = preHead
	return *preHead
}

func (this *MyLinkedList) Get(index int) int {
	if index+1 > this.lenth {
		return -1
	}
	temp := &MyLinkedList{0, this.Next, this.prev, 0}
	for i := 0; i < index+1; i++ {
		temp = temp.Next
	}

	return temp.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	if this.lenth == 0 { //说明第一次加元素
		newHead := &MyLinkedList{val, this.prev, this, 0} //新头节点头指向虚拟头，尾部指向pre的next
		//尾巴指向 这个头
		this.prev.prev = newHead
		this.Next = newHead
		this.lenth++
	} else {
		newHead := &MyLinkedList{val, this.Next, this, 0}
		this.Next.prev = newHead
		this.Next = newHead
		this.lenth++
	}
}

func (this *MyLinkedList) AddAtTail(val int) {
	if this.lenth == 0 {
		newTail := &MyLinkedList{val, this.prev, this, 0}
		this.Next = newTail
		this.prev.prev = newTail
		this.lenth++
	} else {
		newTail := &MyLinkedList{val, this.prev, this.prev.prev, 0}
		this.prev.prev.Next = newTail
		this.prev.prev = newTail
		this.lenth++
	}
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if this.lenth < index {
		return
	}
	if this.lenth == index {
		if index == 0 { //和第一次加元素一样的情况
			newHead := &MyLinkedList{val, this.prev, this, 0} //新头节点头指向虚拟头，尾部指向pre的next
			//尾巴指向 这个头
			this.prev.prev = newHead
			this.Next = newHead
			this.lenth++
		} else {
			newTail := &MyLinkedList{val, this.prev, this.prev.prev, 0}
			this.prev.prev.Next = newTail
			this.prev.prev = newTail
			this.lenth++
		}
	} else {
		pointer := &MyLinkedList{val, this.Next, nil, 0}
		for i := 0; i < index+1; i++ {
			pointer = pointer.Next
		}
		newNode := &MyLinkedList{val, pointer, pointer.prev, 0}
		pointer.prev.Next = newNode
		pointer.prev = newNode
		this.lenth++
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if this.lenth < index+1 {
		return
	} else {
		pointer := &MyLinkedList{0, this.Next, nil, 0}
		for i := 0; i < index+1; i++ {
			pointer = pointer.Next
		}
		pointer.prev.Next = pointer.Next
		pointer.Next.prev = pointer.prev
		pointer.Next = nil
		pointer.prev = nil
		this.lenth--
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
