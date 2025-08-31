package main

import (
	"errors"
	"fmt"
)

// Очередь, где N - количество элементов в очереди: 

// Сложность вставки элемента: O(1)

// Сложность удаления элемента: O(1)

// Задание: реализовать очередь на базе кольцевого буфера,
// связного списка, двух стеков. Когда мы говорим про очередь,
// мы ожидаем, что у структуры будут следующие методы:

// 1) push(item) — добавляет элемент в конец очереди;

// 2) pop() — берёт элемент из начала очереди и удаляет его;

// 3) peek() — берёт элемент из начала очереди без удаления;

// 4) size() — возвращает количество элементов в очереди.

type CircleQueue struct {
	queue     []interface{} // массив элементов очереди
	head      int           // индекс, по которому нужно извлекать элемент, если очередь не пустая;
	tail      int           // индекс, по которому нужно добавлять элемент, если в очереди есть место;
	max_n     int           // максимально возможное количество элементов в очереди;
	queueSize int           // размер очереди;
}

func NewCircleQueue(capacity int) *CircleQueue {
	return &CircleQueue{
		queue:     make([]interface{}, capacity),
		max_n:     capacity,
		head:      0,
		tail:      0,
		queueSize: 0,
	}
}

func (q *CircleQueue) push(item interface{}) error {
	if q.queueSize < q.max_n {
		q.queue[q.tail] = item
		q.tail = (q.tail + 1) % q.max_n
		q.queueSize++
		return nil
	}

	return errors.New("Queue is full")
}

func (q *CircleQueue) pop() interface{} {
	if q.queueSize > 0 {
		value := q.queue[q.head]
		q.queueSize--
		q.head = (q.head + 1) % q.max_n
		return value
	}
	return nil
}

func (q *CircleQueue) peek() interface{} {
	if q.queueSize > 0 {
		value := q.queue[q.head]
		return value
	}
	return nil
}

func (q *CircleQueue) size() int {
	return q.queueSize
}

func main() {
	q := NewCircleQueue(10)
	q.push(1)
	q.push(2)
	q.push(3)
	q.push(4)
	q.push(5)

	fmt.Println(q.pop())
}
