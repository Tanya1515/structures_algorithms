package main

//OK
// Dec - очередь, которая позволяет доставать/добавлять элементы в начала и конец, 
// где N - количество элементов в односвязном списке: 

// Сложность вставки элемента: O(1) (вне зависимости от того, куда добавляется элемент: 
// в начало или в конец)

// Сложность удаления элемента: O(1) (вне зависимости от того, куда добавляется элемент: 
// в начало или в конец)

// Задание: ревлизовать дек. Интерфейс дека подразумевает,
// что в нём будут реализованы следующие методы:

// 1) push_back(item) — вставка нового элемента в конец;

// 2) pop_back() — удаление последнего элемента;

// 3) push_front(item) — вставка нового элемента в начало;

// 4) pop_front() — удаление первого элемента;

// 5) size() — количество элементов в очереди.

type Elem struct {
	data interface{}
	prev *Elem
	next *Elem
}

type Dec struct {
	head    *Elem
	tail    *Elem
	decSize int
}

func (d *Dec) push_back(item interface{}) {
	d.decSize++
	newElem := &Elem{data: item, prev: d.tail, next: nil}
	if d.decSize == 0 {
		d.tail = newElem
		d.head = d.tail
		return
	}
	d.tail.next = newElem
	d.tail = newElem
}

func (d *Dec) pop_back() interface{} {
	if d.decSize == 0 {
		return nil
	}
	value := d.tail.data
	d.decSize--
	oldTail := d.tail
	d.tail = d.tail.prev
	oldTail.prev = nil
	if d.tail != nil {
		d.tail.next = nil
		return value
	}
	d.head = nil
	return value
}

func (d *Dec) push_front(item interface{}) {
	d.decSize++
	newElem := &Elem{data: item, prev: nil, next: d.head}
	if d.decSize == 0 {
		d.head = newElem
		d.tail = d.head
		return
	}
	d.head.prev = newElem
	d.head = newElem
}

func (d *Dec) pop_front() interface{} {
	if d.decSize == 0 {
		return nil
	}
	d.decSize--
	value := d.head.data
	oldHead := d.head
	d.head = d.head.next
	oldHead.next = nil
	if d.head != nil {
		d.head.prev = nil
		return value
	}
	d.tail = nil
	return value
}

func (d *Dec) size() int {
	return d.decSize
}
