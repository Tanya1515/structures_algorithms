package main

// Задание: реализовать очередь на базе кольцевого буфера,
// связного списка, двух стеков. Когда мы говорим про очередь,
// мы ожидаем, что у структуры будут следующие методы:

// 1) push(item) — добавляет элемент в конец очереди;

// 2) pop() — берёт элемент из начала очереди и удаляет его;

// 3) peek() — берёт элемент из начала очереди без удаления;

// 4) size() — возвращает количество элементов в очереди.

type Node struct {
	next  *Node
	value interface{}
}

type ListQueue struct {
	head      *Node
	tail      *Node
	queueSize int
}

func (q *ListQueue) push(item interface{}) {
	if q.queueSize == 0 {
		newElem := &Node{next: nil, value: item}
		q.head = newElem
		q.tail = newElem
		q.queueSize++
		return
	}
	newElem := &Node{next: nil, value: item}
	q.tail.next = newElem
	q.tail = newElem
	q.queueSize++
	return
}

func (q *ListQueue) pop() interface{} {
	if q.queueSize == 0 {
		return nil
	}
	value := q.head.value
	oldHead := q.head
	q.head = q.head.next
	oldHead.next = nil
	q.queueSize--
	if q.queueSize == 0 {
		q.tail = nil
	}
	return value
}

func (q *ListQueue) peek() interface{} {
	if q.queueSize == 0 {
		return nil
	}
	value := q.head.value
	return value
}

func (q *ListQueue) size() int {
	return q.queueSize
}
