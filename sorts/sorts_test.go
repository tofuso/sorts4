package sorts

import (
	"testing"
)

var (
	source = []int{89, 94, 34, 78, 6, 19, 4, 95, 62, 79}
	correct = []int{4, 6, 19, 34, 62, 78, 79, 89, 94, 95}
)

func TestInsertionSort(t *testing.T) {
	numbers := make([]int, len(source))
	copy(numbers, source)
	InsertionSort(numbers)
	for i, num := range numbers {
		if num != correct[i] {
			t.Fatal("not correct: ", numbers, correct)
		}
	}
}

func TestMergeSort(t *testing.T) {
	numbers := make([]int, len(source))
	copy(numbers, source)
	MergeSort(numbers)
	for i, num := range numbers {
		if num != correct[i] {
			t.Fatal("not correct: ", numbers, correct)
		}
	}
}

func TestHeapSort(t *testing.T){
	numbers := make([]int, len(source))
	copy(numbers, source)
	HeapSort(numbers)
	for i, num := range numbers {
		if num != correct[i] {
			t.Fatal("not correct: ", numbers, correct)
		}
	}
}

func TestQuickSort(t *testing.T){
	numbers := make([]int, len(source))
	copy(numbers, source)
	QuickSort(numbers)
	for i, num := range numbers {
		if num != correct[i] {
			t.Fatal("not correct: ", numbers, correct)
		}
	}
}