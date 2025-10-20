package main

/*

MinHeap

Свойство: Каждый родитель ≤ своих потомков. Корень = минимум.

Добавление элемента в min-heap: 

Time Complexity: O(log N)

1) Добавить в конец: Помещаем новый элемент в конец массива (последняя позиция на последнем уровне)

2) Всплытие (Sift-Up):

	- Сравниваем добавленный элемент с его родителем parent_index = (index - 1) // 2

	- Если элемент меньше родителя → меняем их местами

	- Повторяем с новым родителем, пока не дойдем до корня или пока элемент не станет ≥ родителя

3) Результат: Элемент занимает правильную позицию, свойство кучи восстановлено. 



Удаление корня (минимума) из min-heap:

Time Complexity: O(log N)

1) Сохранить корень: Запоминаем значение корневого элемента (это минимум)

2) Заменить корень последним элементом: Берем последний элемент кучи и перемещаем его в корень

3) Удалить последний элемент: Уменьшаем размер кучи на 1

4) Просеивание вниз (Sift-Down):

	- Сравниваем новый корень с его потомками left_child = 2*i + 1, right_child = 2*i + 2

	- Находим наименьший из трех (родитель + два потомка)

	- Если наименьший - не родитель → меняем родителя с наименьшим потомком

	- Повторяем для новой позиции, пока элемент не станет ≤ обоих потомков или не достигнет листа

*/


// Просеивание вверх для min-heap (аналогично вашему стилю)
func siftUpMin(heap []int, index int) []int {
    current := index
    
    for current > 0 {
        parentIndex := (current - 1) / 2
        
        // Для min-heap: если родитель МЕНЬШЕ или равен текущему - свойство выполнено
        if heap[parentIndex] <= heap[current] {
            break
        }
        
        // Меняем местами с родителем (родитель больше - нарушает min-heap)
        heap[parentIndex], heap[current] = heap[current], heap[parentIndex]
        current = parentIndex
    }
    
    return heap
}

// Просеивание вниз для min-heap
func siftDownMin(heap []int, index int) []int {
    heapSize := len(heap)
    current := index
    
    for {
        smallest := current
        left := 2*current + 1
        right := 2*current + 2
        
        // Ищем наименьший среди текущего и левого потомка
        if left < heapSize && heap[left] < heap[smallest] {
            smallest = left
        }
        
        // Ищем наименьший среди текущего наименьшего и правого потомка
        if right < heapSize && heap[right] < heap[smallest] {
            smallest = right
        }
        
        // Если наименьший - текущий, то свойство min-heap выполнено
        if smallest == current {
            break
        }
        
        // Меняем с наименьшим потомком
        heap[current], heap[smallest] = heap[smallest], heap[current]
        current = smallest
    }
    
    return heap
}

// Вставка в min-heap (аналог вашей функции)
func InsertMin(heap []int, value int) []int {
    heap = append(heap, value)
    heap = siftUpMin(heap, len(heap)-1)
    return heap
}

// Удаление минимального элемента (корня) из min-heap
func ExtractMin(heap []int) (int, []int) {
    if len(heap) == 0 {
        return -1, heap
    }
    
    // Минимальный элемент - корень
    min := heap[0]
    
    // Перемещаем последний элемент в корень
    heap[0] = heap[len(heap)-1]
    heap = heap[:len(heap)-1]
    
    // Просеиваем новый корень вниз
    if len(heap) > 0 {
        heap = siftDownMin(heap, 0)
    }
    
    return min, heap
}