package main

// Сложность по времени: O(logN)

// Сложность по памяти: O(logN), поскольку на каждом шаге рекурсии будет создаваться новый фрейм стека

// Задание: Необходимо найти элемент в отсортированном массиве.
// В алгоритме бинарного поиска мы будем проверять центральный
// элемент отсортированного массива.

// Чтобы узнать индекс середины массива, необходимо сложить индексы правой
// и левой границ и полелить пополам: mid = (left + right) // 2

// Если центральный элемент:

// 1) равен искомому, то вернём его индекс;

// 2) больше искомого, то продолжим рекурсивный поиск в левой половине массива;

// 3) меньше искомого, то продолжим рекурсивный поиск в правой половине массива.

// Необходимо решить задачу для двух вариантов: когда на вход поступает массив,
// отсортированный по возрастанию и по убыванию.

func recursiveBinSearchInIncreaseArray(slice []int, elem, leftInd, rightInd int) int {
	if rightInd < leftInd {
		return -1
	}
	mid := (leftInd + rightInd) / 2

	if elem == slice[mid] {
		return mid
	}

	if elem > slice[mid] {
		return recursiveBinSearchInIncreaseArray(slice, elem, mid+1, rightInd)
	} else {
		return recursiveBinSearchInIncreaseArray(slice, elem, leftInd, mid-1)
	}
}

func recursiveBinSearchInDecreaseArray(slice []int, elem, leftInd, rightInd int) int {
	if rightInd < leftInd {
		return -1
	}
	mid := (leftInd + rightInd) / 2

	if elem == slice[mid] {
		return mid
	}

	if elem > slice[mid] {
		return recursiveBinSearchInDecreaseArray(slice, elem, leftInd, mid-1)
	} else {
		return recursiveBinSearchInDecreaseArray(slice, elem, mid+1, rightInd)
	}
}
