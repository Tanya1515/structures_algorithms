package main

import "fmt"

// Задание:
// Реализовать стек на golang на базе массива. Должны быть реализованы следующие методы:

// 1) push(item) — добавляет элемент на вершину стека;

// 2) pop() — возвращает элемент с вершины стека и удаляет его;

// 3) size() — возвращает размер стека (количество лежащих в нём элементов);

// 4) peek() или top() — возвращает элемент с вершины стека, не удаляя его;

// 5) isEmpty() — определяет, пуст ли стек.

type Stack struct {
	array []interface{}
}

func (s *Stack) print() {
	for _, value := range s.array {
		fmt.Println(value)
	}
}

func (s *Stack) push(value interface{}) {
	s.array = append(s.array, value)
}

func (s *Stack) pop() interface{} {
	elem := s.array[len(s.array)-1]
	s.array = s.array[:len(s.array)-1]
	return elem
}

func (s *Stack) size() int {
	return len(s.array)
}

func (s *Stack) top() interface{} {
	return s.array[len(s.array)-1]
}

func (s *Stack) isEmpty() bool {
	if len(s.array) == 0 {
		return true
	}
	return false
}

// func main() {
// 	s := new(Stack)
// 	array := make([]interface{}, 0, 10)
// 	s.array = array

// 	s.push("1")
// 	s.push("2")
// 	s.push("3")
// 	s.push("4")
// 	s.push("5")
// 	s.push("6")
// 	s.push("7")

// 	fmt.Println()
// 	fmt.Println("Stack: ")
// 	s.print()

// 	fmt.Println("Popped element: ", s.pop())
// 	fmt.Println("Stack size: ", s.size())
// 	fmt.Println("Get element without delete operation: ", s.top())
// 	fmt.Println("Stack: ")
// 	s.print()

// 	fmt.Println("Is stack empty: ", s.isEmpty())
// }
