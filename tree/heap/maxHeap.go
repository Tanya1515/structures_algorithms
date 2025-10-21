package main

/* 

MaxHeap

Time Complexity: O(log N)

Свойство: Каждый родитель ≥ своих потомков. Корень = максимум.

Добавление элемента в max-heap:

1) Добавить в конец: Помещаем новый элемент в конец массива

2) Всплытие (Sift-Up): 

	- Сравниваем добавленный элемент с его родителем parent_index = (index - 1) // 2

	- Если элемент больше родителя → меняем их местами

	- Повторяем с новым родителем, пока не дойдем до корня или пока элемент не станет ≤ родителя

3) Результат: Элемент занимает правильную позицию, свойство кучи восстановлено.



Удаление корня (максимума) из max-heap:

Time Complexity: O(log N)

1) Сохранить корень: Запоминаем значение корневого элемента (это максимум)

2) Заменить корень последним элементом: Берем последний элемент кучи и перемещаем его в корень

3) Удалить последний элемент: Уменьшаем размер кучи на 1

4) Просеивание вниз (Sift-Down):

	- Сравниваем новый корень с его потомками left_child = 2*i + 1, right_child = 2*i + 2

	- Находим наибольший из трех (родитель + два потомка)

	- Если наибольший - не родитель → меняем родителя с наибольшим потомком

	- Повторяем для новой позиции, пока элемент не станет ≥ обоих потомков или не достигнет листа

*/


// Insert добавляет новый элемент в max-heap
func Insert(arr []int, value int) []int {
    arr = append(arr, value)
    siftUp(arr, len(arr) - 1)
	return arr
}

// siftUp поднимает элемент вверх для восстановления свойства max-heap
func siftUp(heap []int, index int) {
    
    for index  > 0 {
        parent := (index - 1) / 2
        
        // Проверяем свойство max-heap
        if heap[parent] >= heap[index] {
            break
        }
        
        // Обмен с родителем
        heap[parent], heap[index] = heap[index], heap[parent]
        index = parent
    }
}


// Находит индекс максимального элемента среди родителя и двух потомков
func findMaxIndex(heap []int, i, heapSize int) int {
    maxIndex := i
    left := 2*i + 1
    right := 2*i + 2
    
    if left < heapSize && heap[left] > heap[maxIndex] {
        maxIndex = left
    }
    
    if right < heapSize && heap[right] > heap[maxIndex] {
        maxIndex = right
    }
    
    return maxIndex
}

// Просеивание элемента вниз (heapify down)
func siftDown(heap []int, i int) []int {
    heapSize := len(heap)
    current := i
    
    for {
        maxIndex := findMaxIndex(heap, current, heapSize)
        
        if maxIndex == current {
            break
        }
        
        heap[current], heap[maxIndex] = heap[maxIndex], heap[current]
        current = maxIndex
    }
    
    return heap
}

// Удаление корня из max-heap (корректная версия)
func ExtractMax(heap []int) (int, []int) {
    if len(heap) == 0 {
        return -1, heap
    }
    
    // Сохраняем максимальный элемент (корень)
    max := heap[0]
    
    // Перемещаем последний элемент в корень
    heap[0] = heap[len(heap)-1]
    heap = heap[:len(heap)-1]
    
    // Просеиваем новый корень вниз
    if len(heap) > 0 {
        heap = siftDown(heap, 0)
    }
    
    return max, heap
}