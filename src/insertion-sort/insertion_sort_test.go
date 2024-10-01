package insertion_sort

import (
	"reflect"
	"testing"
)

func TestInsertionSortWithEmptyArray(t *testing.T) {
	sortedArray := InsertionSort([]int{})
	var wanted []int

	if len(sortedArray) != len(wanted) {
		t.Errorf("InsertionSort() = %v; want %v", sortedArray, wanted)
	}
}

func TestInsertionSortWithOneElement(t *testing.T) {
	sortedArray := InsertionSort([]int{1})
	wanted := []int{1}

	if !reflect.DeepEqual(sortedArray, wanted) {
		t.Errorf("InsertionSort() = %v; want %v", sortedArray, wanted)
	}
}

func TestInsertionSortWithFiveElements(t *testing.T) {
	sortedArray := InsertionSort([]int{5, 3, 4, 1, 2})
	wanted := []int{1, 2, 3, 4, 5}

	if !reflect.DeepEqual(sortedArray, wanted) {
		t.Errorf("InsertionSort() = %v; want %v", sortedArray, wanted)
	}
}
