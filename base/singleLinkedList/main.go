package main

import "fmt"

// Задание:
// Реалзовать односвзяный список, а именно:

// 1) Функция печати односвязного списка

// 2) Функция добавления элемента в список

// 3) Функция удаления элемента из списка

type Node struct {
	value interface{}
	next  *Node
}

func PrintList(list *Node) {
	var i int
	for list != nil {
		fmt.Printf("Node %d: %v\n", i, list.value)
		i++
		list = list.next
	}
}

func AddElem(head *Node, index int, value interface{}) *Node {
	if index == 0 {
		return &Node{value: value, next: head}
	}
	tmp := head
	i := 0
	for tmp != nil {
		if i == index-1 {
			newElem := Node{value: value, next: tmp.next}
			tmp.next = &newElem
			break
		}
		i++
		tmp = tmp.next
	}

	return head
}

func DeleteElem(head *Node, index int) *Node {

	if index == 0 {
		return head.next
	}
	tmp := head
	prev := head
	i := 0

	for tmp != nil {
		if i == index {
			prev.next = tmp.next
			break
		}
		i++
		prev = tmp
		tmp = tmp.next
	}

	return head
}

func main() {
	fifthElem := Node{value: "fifth", next: nil}
	forthElem := Node{value: "forth", next: &fifthElem}
	thirdElem := Node{value: "third", next: &forthElem}
	secondElem := Node{value: "second", next: &thirdElem}
	firstElem := Node{value: "first", next: &secondElem}

	fmt.Println("Base list: ")
	PrintList(&firstElem)

	head := AddElem(&firstElem, 0, "CheckFirst")

	fmt.Println()
	fmt.Println("Add new first Elem: ")
	PrintList(head)

	fmt.Println()
	fmt.Println("Add new center Elem: ")
	head = AddElem(head, 2, "CheckCenter")

	PrintList(head)

	fmt.Println()
	fmt.Println("Add new last Elem: ")
	head = AddElem(head, 7, "CheckLast")
	PrintList(head)

	fmt.Println()
	head = DeleteElem(head, 0)
	fmt.Println("Remove first Elem: ")
	PrintList(head)

	fmt.Println()
	fmt.Println("Remove center Elem: ")
	head = DeleteElem(head, 1)

	PrintList(head)

	fmt.Println()
	fmt.Println("Remove last Elem: ")
	head = DeleteElem(head, 5)
	PrintList(head)

}
