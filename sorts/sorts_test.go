package sorts

import (
	"testing"
)

var (
	numbers = []int{89, 94, 34, 78, 6, 19, 4, 95, 62, 79}
	correct = []int{4, 6, 19, 34, 62, 78, 79, 89, 94, 95}
)

func TestInsertionSort(t *testing.T) {
	InsertionSort(numbers)
	for i, num := range numbers {
		if num != correct[i] {
			t.Fatal("not correct: ", numbers, correct)
		}
	}
}

func TestMergeSort(t *testing.T) {
	MergeSort(numbers)
	for i, num := range numbers {
		if num != correct[i] {
			t.Fatal("not correct: ", numbers, correct)
		}
	}
}

func TestHeapSort(t *testing.T){
	HeapSort(numbers)
	for i, num := range numbers {
		if num != correct[i] {
			t.Fatal("not correct: ", numbers, correct)
		}
	}
}