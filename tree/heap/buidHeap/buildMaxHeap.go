package main

/*

Time Complexity: O(N)

Space Complexity: O(1)

*/

// heapify преобразует поддерево в max-heap (итеративная версия)
func heapifyMax(arr []int, n, i int) {
	current := i
	for {
		largest := current     // Инициализируем наибольший элемент как корень
		left := 2*current + 1  // Левый потомок
		right := 2*current + 2 // Правый потомок

		// Если левый потомок существует и больше текущего наибольшего
		if left < n && arr[left] > arr[largest] {
			largest = left
		}

		// Если правый потомок существует и больше текущего наибольшего
		if right < n && arr[right] > arr[largest] {
			largest = right
		}

		// Если наибольший элемент не корень
		if largest == current {
			break
		}

		// Меняем местами и продолжаем
		arr[current], arr[largest] = arr[largest], arr[current]
		current = largest
	}
}

// heapSort выполняет пирамидальную сортировку
func buildMaxHeap(arr []int) {
	n := len(arr)

	// Построение max-heap (перегруппировка массива)
	for i := n/2 - 1; i >= 0; i-- {
		heapifyMax(arr, n, i)
	}
}
