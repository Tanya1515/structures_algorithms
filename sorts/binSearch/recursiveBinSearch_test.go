package main

import "testing"

func TestAbsentElementsRec(t *testing.T) {
	slice := []int{2, 4, 6, 8, 10}
	tests := []struct {
		elem     int
		leftInd  int
		rightInd int
		expected int
	}{
		{1, 0, 4, -1},  // Меньше минимального
		{11, 0, 4, -1}, // Больше максимального
		{5, 0, 4, -1},  // Между элементами
		{3, 0, 4, -1},  // Между элементами
		{7, 0, 4, -1},  // Между элементами
	}

	for _, test := range tests {
		result := recursiveBinSearchInIncreaseArray(slice, test.elem, test.leftInd, test.rightInd)
		if result != test.expected {
			t.Errorf("Для elem=%d ожидалось %d, получено %d", test.elem, test.expected, result)
		}
	}
}

func TestBasicCasesRec(t *testing.T) {
	tests := []struct {
		slice    []int
		elem     int
		leftInd  int
		rightInd int
		expected int
	}{
		// Элемент в середине
		{[]int{1, 3, 5, 7, 9, 11}, 5, 0, 5, 2},
		// Первый элемент
		{[]int{1, 3, 5, 7, 9, 11}, 1, 0, 5, 0},
		// Последний элемент
		{[]int{1, 3, 5, 7, 9, 11}, 11, 0, 5, 5},
		// Элемент в правой половине
		{[]int{1, 3, 5, 7, 9, 11}, 7, 0, 5, 3},
		// Элемент в левой половине
		{[]int{1, 3, 5, 7, 9, 11}, 3, 0, 5, 1},
	}

	for _, test := range tests {
		result := recursiveBinSearchInIncreaseArray(test.slice, test.elem, test.leftInd, test.rightInd)
		if result != test.expected {
			t.Errorf("Для slice=%v, elem=%d, left=%d, right=%d ожидалось %d, получено %d",
				test.slice, test.elem, test.leftInd, test.rightInd, test.expected, result)
		}
	}
}


func TestAbsentElementsDec(t *testing.T) {
	slice := []int{10, 8, 6, 4, 2}
	tests := []struct {
		elem     int
		leftInd  int
		rightInd int
		expected int
	}{
		{1, 0, 4, -1},  // Меньше минимального
		{11, 0, 4, -1}, // Больше максимального
		{5, 0, 4, -1},  // Между элементами
		{3, 0, 4, -1},  // Между элементами
		{7, 0, 4, -1},  // Между элементами
	}

	for _, test := range tests {
		result := recursiveBinSearchInDecreaseArray(slice, test.elem, test.leftInd, test.rightInd)
		if result != test.expected {
			t.Errorf("Для elem=%d ожидалось %d, получено %d", test.elem, test.expected, result)
		}
	}
}

func TestBasicCasesDec(t *testing.T) {
	tests := []struct {
		slice    []int
		elem     int
		leftInd  int
		rightInd int
		expected int
	}{
		// Элемент в середине
		{[]int{11, 9, 7, 5, 3, 1}, 5, 0, 5, 3},
		// Первый элемент
		{[]int{11, 9, 7, 5, 3, 1}, 1, 0, 5, 5},
		// Последний элемент
		{[]int{11, 9, 7, 5, 3, 1}, 11, 0, 5, 0},
		// Элемент в правой половине
		{[]int{11, 9, 7, 5, 3, 1}, 7, 0, 5, 2},
		// Элемент в левой половине
		{[]int{11, 9, 7, 5, 3, 1}, 3, 0, 5, 4},
	}

	for _, test := range tests {
		result := recursiveBinSearchInDecreaseArray(test.slice, test.elem, test.leftInd, test.rightInd)
		if result != test.expected {
			t.Errorf("Для slice=%v, elem=%d, left=%d, right=%d ожидалось %d, получено %d",
				test.slice, test.elem, test.leftInd, test.rightInd, test.expected, result)
		}
	}
}
