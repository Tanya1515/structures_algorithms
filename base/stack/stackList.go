package main

import "fmt"

// Задание:
// Реализовать стек на golang на базе списка. Должны быть реализованы следующие методы:

// 1) push(item) — добавляет элемент на вершину стека;

// 2) pop() — возвращает элемент с вершины стека и удаляет его;

// 3) size() — возвращает размер стека (количество лежащих в нём элементов);

// 4) peek() или top() — возвращает элемент с вершины стека, не удаляя его;

// 5) isEmpty() — определяет, пуст ли стек.

type Node struct {
	next  *Node
	value interface{}
}

type StackList struct {
	head      *Node
	stackSize int
}

func (s *StackList) print() {
	tmp := s.head
	var i int
	for tmp != nil {
		fmt.Printf("Element %d: %v\n", i, tmp.value)
		i++
		tmp = tmp.next
	}
}

func (s *StackList) push(value interface{}) {
	newElem := &Node{next: s.head, value: value}

	s.head = newElem
	s.stackSize++
}

func (s *StackList) pop() interface{} {
	if s.stackSize == 0 {
		return nil
	}
	value := s.head.value
	s.head = s.head.next
	s.stackSize--
	return value
}

func (s *StackList) size() int {
	return s.stackSize
}

func (s *StackList) top() interface{} {
	if s.isEmpty() {
		return nil
	}
	return s.head.value
}

func (s *StackList) isEmpty() bool {
	if s.stackSize == 0 {
		return true
	}
	return false
}

func main() {
	s := new(StackList)

	s.push("1")
	s.push("2")
	s.push("3")
	s.push("4")
	s.push("5")
	s.push("6")
	s.push("7")

	fmt.Println()
	fmt.Println("Stack: ")
	s.print()

	fmt.Println("Popped element: ", s.pop())
	fmt.Println("Stack size: ", s.size())
	fmt.Println("Get element without delete operation: ", s.top())
	fmt.Println("Stack: ")
	s.print()

	fmt.Println("Is stack empty: ", s.isEmpty())
}
