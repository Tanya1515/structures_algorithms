package main

import (
	"errors"
	"fmt"
)

//OK
// Двусвязный список, где N - количество элементов в двусявязном списке:

// Сложность вставки элемента: O(1) (в любое место списка,
// если известно после/перед каким элементом необходимо добавить новый).

// Сложность удаления элемента: O(1)

// Задание:
// Реалзовать односвзяный список, а именно:

// 1) Функция печати односвязного списка

// 2) Функция добавления элемента в список

// 3) Функция удаления элемента из списка

type Node struct {
	data interface{}
	prev *Node
	next *Node
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
	size int
}

func (l *DoublyLinkedList) PrintList() {
	var i int
	tmp := l.head
	for tmp != nil {
		fmt.Printf("Node %d: %v\n", i, tmp.data)
		i++
		tmp = tmp.next
	}
}

func (l *DoublyLinkedList) InsertAfter(node *Node, data interface{}) *Node {
	if node == nil {
		return nil
	}
	newNode := Node{data: data, prev: node, next: node.next}

	if node.next == nil {
		l.tail = &newNode
	} else {
		node.next.prev = &newNode
	}

	node.next = &newNode

	l.size++
	return &newNode
}

func (l *DoublyLinkedList) InsertBefore(node *Node, data interface{}) *Node {
	if node == nil {
		return nil
	}
	newNode := Node{data: data, prev: node.prev, next: node}

	if node.prev == nil {
		l.head = &newNode
	} else {
		node.prev.next = &newNode
	}
	node.prev = &newNode

	l.size++

	return &newNode
}

func (l *DoublyLinkedList) DeleteElem(node *Node) error {
	if node == nil {
		return errors.New("Incorrect node recieved")
	}

	defer func() {
		l.size--
	}()
	if node.next == nil {
		l.tail = node.prev
		if node.prev != nil {
			node.prev.next = nil
		}
		return nil
	}

	if node.prev == nil {
		l.head = node.next
		node.next.prev = nil
		return nil
	}

	node.prev.next = node.next
	node.next.prev = node.prev

	return nil
}

func main() {
	firstNode := Node{data: "first"}
	secondNode := Node{data: "second"}
	thirdNode := Node{data: "third"}
	forthNode := Node{data: "forth"}

	firstNode.next = &secondNode
	secondNode.prev = &firstNode
	secondNode.next = &thirdNode
	thirdNode.prev = &secondNode
	thirdNode.next = &forthNode
	forthNode.prev = &thirdNode

	l := DoublyLinkedList{head: &firstNode, tail: &forthNode, size: 4}
	l.PrintList()

	fmt.Println()
	fmt.Println("Modified list")
	fmt.Println()

	// l.DeleteElem(&thirdNode)
	// l.PrintList()

	l.InsertAfter(&forthNode, "checkLast")
	l.PrintList()

	fmt.Println("Modified list")
	fmt.Println()

	l.InsertBefore(&firstNode, "moveHead")
	l.PrintList()

}
