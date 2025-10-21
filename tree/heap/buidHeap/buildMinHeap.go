package main

/*

Time Complexity: O(N)

Space Complexity: O(1)

*/

// heapifyMin преобразует поддерево в min-heap (итеративная версия)
func heapifyMin(arr []int, n, i int) {
    current := i
    for {
        smallest := current     // Инициализируем наименьший элемент как корень
        left := 2*current + 1   // Левый потомок
        right := 2*current + 2  // Правый потомок

        // Если левый потомок существует и МЕНЬШЕ текущего наименьшего
        if left < n && arr[left] < arr[smallest] {
            smallest = left
        }

        // Если правый потомок существует и МЕНЬШЕ текущего наименьшего
        if right < n && arr[right] < arr[smallest] {
            smallest = right
        }

        // Если наименьший элемент уже корень - выходим
        if smallest == current {
            break
        }

        // Меняем местами и продолжаем
        arr[current], arr[smallest] = arr[smallest], arr[current]
        current = smallest
    }
}

// buildHeapMin строит min-heap из массива
func buildHeapMin(arr []int) {
    n := len(arr)
    
    // Построение min-heap (перегруппировка массива)
    for i := n/2 - 1; i >= 0; i-- {
        heapifyMin(arr, n, i)
    }
}