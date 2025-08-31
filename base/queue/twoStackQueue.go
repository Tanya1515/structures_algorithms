package main

// Очередь, где N - количество элементов в очереди: 

// Сложность вставки элемента: O(1) 

// Сложность удаления элемента: амортизированная O(1)

// Задание: реализовать очередь на базе кольцевого буфера,
// связного списка, двух стеков. Когда мы говорим про очередь,
// мы ожидаем, что у структуры будут следующие методы:

// 1) push(item) — добавляет элемент в конец очереди;

// 2) pop() — берёт элемент из начала очереди и удаляет его;

// 3) peek() — берёт элемент из начала очереди без удаления;

// 4) size() — возвращает количество элементов в очереди.

// Основная идея: есть два массива. Первый массив представляет 
// собой первый стек для добавления элементов. Второй массив 
// представлеят собой второй стек, откуда элементы будут доставаться. 
// Если элементы во втором стэке закончились, то добавляем их 
// из первого стека во второй. Причем первый стек начинает 
// обходится с последнего элемента, но они попадают в начало второго стека. 
// Из-за переливания данных сложность операции pop() или peek() 
// - амортизированная константа. 

type StackQueue struct {
	addStack []interface{}
	getStack []interface{}
}

func (q *StackQueue) transfer() {
	if len(q.getStack) == 0 {
		for i := len(q.addStack) - 1; i >= 0; i-- {
			q.getStack = append(q.getStack, q.addStack[i])
		}
		q.addStack = q.addStack[:0]
	}

}

func (q *StackQueue) push(item interface{}) {
	q.addStack = append(q.addStack, item)
	return
}

func (q *StackQueue) pop() interface{} {
	q.transfer()
	if len(q.getStack) > 0 {
		elem := q.getStack[len(q.getStack)-1]
		q.getStack = q.getStack[:len(q.getStack)-1]
		return elem
	}
	return nil
}

func (q *StackQueue) peek() interface{} {
	q.transfer()
	if len(q.getStack) > 0 {
		elem := q.getStack[len(q.getStack)-1]
		return elem
	}
	return nil
}

func (q *StackQueue) size() int {
	return len(q.addStack) + len(q.getStack)
}
