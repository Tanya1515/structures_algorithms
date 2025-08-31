package main

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		// Базовые случаи
		{
			name:     "обычный несортированный массив",
			input:    []int{5, 2, 8, 1, 9},
			expected: []int{1, 2, 5, 8, 9},
		},
		{
			name:     "массив с отрицательными числами",
			input:    []int{-3, 5, -1, 0, 2},
			expected: []int{-3, -1, 0, 2, 5},
		},
		{
			name:     "уже отсортированный массив",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "обратно отсортированный массив",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},

		// Пограничные случаи
		{
			name:     "пустой массив",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "массив из одного элемента",
			input:    []int{42},
			expected: []int{42},
		},
		{
			name:     "массив из двух элементов",
			input:    []int{2, 1},
			expected: []int{1, 2},
		},

		// Случаи с дубликатами
		{
			name:     "массив с повторяющимися элементами",
			input:    []int{3, 1, 4, 1, 5, 9, 2, 6, 5},
			expected: []int{1, 1, 2, 3, 4, 5, 5, 6, 9},
		},
		{
			name:     "все элементы одинаковые",
			input:    []int{7, 7, 7, 7, 7},
			expected: []int{7, 7, 7, 7, 7},
		},

		// Большие массивы
		{
			name:     "большой массив",
			input:    []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
			expected: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := QuickSort(tt.input)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("BubbleSort(%v) = %v, ожидалось %v",
					tt.input, result, tt.expected)
			}

			// Дополнительная проверка: результат должен быть отсортирован
			if !isSorted(result) {
				t.Errorf("BubbleSort(%v) = %v - массив не отсортирован",
					tt.input, result)
			}
		})
	}
}

// Вспомогательная функция для проверки отсортированности
func isSorted(slice []int) bool {
	for i := 1; i < len(slice); i++ {
		if slice[i] < slice[i-1] {
			return false
		}
	}
	return true
}
