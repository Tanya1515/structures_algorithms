package main

import (
	"fmt"
)

// Задание: необходимо реализовать heapSort за О(1) по памяти. 

func sortHeap(heap []int, index int) []int {
	if index == 0 {
		return heap
	}
	parentIndex := (index - 1)/2

	// если реализуем условие для min-heap, то оно меняется: 
	// heap[parentIndex] > heap[index], поскольку на вершине 
	// находится самый большой элемент. 
	if heap[parentIndex] < heap[index] {
		heap[parentIndex], heap[index] = heap[index], heap[parentIndex]
		sortHeap(heap, parentIndex)
	}

	return heap
}

func maxIndex(heap []int, index int) int {
	left := 2*index + 1
	right := 2*index + 2
	maxIndex := index

	heapLen := len(heap)

	if heapLen > left && heap[left] > heap[maxIndex]  {
		maxIndex = left
	}

	if heapLen > right && heap[right] > heap[maxIndex] {
		maxIndex = right
	}

	return maxIndex
}

func sortHeapDown(heap []int, index, indexEnd int) []int {
	if index > indexEnd {
		return heap
	}
	maxIndex := maxIndex(heap[:indexEnd], index)
	if maxIndex != index {
		heap[maxIndex], heap[index] = heap[index], heap[maxIndex]
		heap = sortHeapDown(heap, maxIndex, indexEnd)
	}

	return heap
}

func sortSliceWithHeap(slice []int) []int {
	for i := 0; i <= len(slice) - 1; i++ {
		slice = sortHeap(slice, i)
	}
	
	for i := len(slice); i > 1; i-- {
		slice[0], slice[i-1] = slice[i-1], slice[0]
		slice = sortHeapDown(slice, 0, i-1)
	}

	return slice

}


func main() {
	slice := []int{2, 10, 4, 3, 5, 1}
	slice = sortSliceWithHeap(slice)
	fmt.Println(slice)
}