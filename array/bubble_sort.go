package array

// BubbleSort solves the problem in O(n^2) time and O(1) space.
// wasSwapped is used to track if any swaps occurred in the current pass,
// so we can break early if no swaps were made (i.e., the array is sorted).
//
// This was my solution.
func BubbleSort(input []int) {
	n := len(input)
	for i := 0; i < n; i++ {
		wasSwapped := false
		for j := 0; j < n-i-1; j++ {
			if input[j] > input[j+1] {
				input[j], input[j+1] = input[j+1], input[j]
				wasSwapped = true
			}
		}

		if !wasSwapped {
			break
		}
	}
}

// BubbleSort2 solves the problem in O(n^2) time and O(1) space.
//
// This is the solution from the original repo.
func BubbleSort2(input []int) {
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < len(input); i++ {
			if input[i] < input[i-1] {
				input[i], input[i-1] = input[i-1], input[i]
				swapped = true
			}
		}
	}
}
