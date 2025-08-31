package main

// Сложность по времени: O(logN)

// Сложность по памяти: O(1)

func binSearch(slice []int, elem int) int {
	leftInd := 0
	rightInd := len(slice) - 1

	for leftInd <= rightInd {
		mid := (rightInd + leftInd) / 2
		if elem == slice[mid] {
			return mid
		}
		if elem > slice[mid] {
			leftInd = mid + 1
			continue
		}
		rightInd = mid - 1
	}
	return -1
}
