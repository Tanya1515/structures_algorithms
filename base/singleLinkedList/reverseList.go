package main

type ListNode struct {
	Value interface{}
	Next  *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	current := head // кладем текущий элемент

	for current != nil {
		next := current.Next // кладем следующий элемент
		current.Next = prev  // в Next для текущего записываем предыдущий
		prev = current       // в предыдущий записываем текущий элемент
		current = next       // в текущий записываем следующий
	}

	return prev
}
