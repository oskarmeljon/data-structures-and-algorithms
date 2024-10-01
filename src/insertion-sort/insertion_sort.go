package insertion_sort

/*
InsertionSort sorts an array of integers using the insertion sort algorithm.
The function should take in an array of integers and return an array with the integers sorted in ascending order.
Returned value should be the same array but modified.
*/
func InsertionSort(arr []int) []int {

	// Iterate over the array starting from the second element.
	for i := 1; i < len(arr); i++ {
		// Store the current element in a variable.
		current := arr[i]
		// Store the index of the previous element in a variable.
		j := i - 1

		// Iterate over the array from the current element to the first element.
		for j >= 0 && arr[j] > current {
			// Shift the elements to the right.
			arr[j+1] = arr[j]
			// Move to the previous element.
			j--
		}
		// Insert the current element in the correct position.
		arr[j+1] = current
	}

	return arr
}
