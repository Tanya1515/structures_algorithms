package main

import "math/rand/v2"

// Сложность по времени: O(NlogN)

// худший случай: O(N²) - в этом варианте временная сложность алгоритма 
// зависит от выбора опорного элемента. Такая сложность будет получаться, 
// если сортировать уже отсортированный массив, и каждый раз в качестве 
// опорного элемента выбирать 1ый элемент. 

// Сложность по памяти:  O(N log N), в силу рекурсии создается log N фреймов на стеке, 
// и на каждом фрейме создается несколько массивов, максимальныйй размер которых - N. 
// Однако, есть in-place реализация, при которой не создается новых массивов. В таком 
// случае пространственная сложность - O(logN), которая покрывает размер дополнительных 
// фреймов стека при вызове. 

// Задание: Отсортировать массив при помощи quick sort.

// Алгоритм действий:

// 1) Выбираем опорный элемент.

// 2) Делим массив на 3 подмассива: слева элементы меньше опорного,
// справа — больше, посередине — равные опорному.
// Этот этап часто выделяют в отдельную функцию, которую называют partition.

// 3) Применяем быструю сортировку к двум подмассивам, левому и правому.

// 4) Объединяем результат.

func partition(slice []int, centerInd int) ([]int, []int, []int) {
	centerSlice := make([]int, 0, len(slice))
	left := make([]int, 0, len(slice))
	right := make([]int, 0, len(slice))
	center := slice[centerInd]

	for _, value := range slice {
		if value == center {
			centerSlice = append(centerSlice, value)
			continue
		}
		if value > center {
			right = append(right, value)
			continue
		}
		left = append(left, value)
	}

	return left, centerSlice, right
}

func QuickSort(slice []int) []int {

	if len(slice) < 2 {
		return slice
	}

	centerInd := rand.IntN(len(slice))

	left, center, right := partition(slice, centerInd)

	left = QuickSort(left)
	right = QuickSort(right)

	result := make([]int, 0, len(slice))
	result = append(result, left...)
	result = append(result, center...)
	result = append(result, right...)

	return result
}
