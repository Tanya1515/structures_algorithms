package main

import "testing"

func TestAbsentElements(t *testing.T) {
	slice := []int{2, 4, 6, 8, 10}
	tests := []struct {
		elem     int
		expected int
	}{
		{1, -1},  // Меньше минимального
		{11, -1}, // Больше максимального
		{5, -1},  // Между элементами
		{3, -1},  // Между элементами
		{7, -1},  // Между элементами
	}

	for _, test := range tests {
		result := binSearch(slice, test.elem)
		if result != test.expected {
			t.Errorf("Для elem=%d ожидалось %d, получено %d", test.elem, test.expected, result)
		}
	}
}

func TestBasicCases(t *testing.T) {
	tests := []struct {
		slice    []int
		elem     int
		expected int
	}{
		// Элемент в середине
		{[]int{1, 3, 5, 7, 9, 11}, 5, 2},
		// Первый элемент
		{[]int{1, 3, 5, 7, 9, 11}, 1, 0},
		// Последний элемент
		{[]int{1, 3, 5, 7, 9, 11}, 11, 5},
		// Элемент в правой половине
		{[]int{1, 3, 5, 7, 9, 11}, 7, 3},
		// Элемент в левой половине
		{[]int{1, 3, 5, 7, 9, 11}, 3, 1},
	}

	for _, test := range tests {
		result := binSearch(test.slice, test.elem)
		if result != test.expected {
			t.Errorf("Для slice=%v, elem=%d, ожидалось %d, получено %d",
				test.slice, test.elem, test.expected, result)
		}
	}
}
