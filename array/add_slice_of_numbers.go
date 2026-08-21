package array

// AddSliceOfTwoNumbers solves the problem in O(n) time and O(n) space.
//
// This is my answer, which I think is more readable than the AI solution.
func AddSliceOfTwoNumbers(num1, num2 []int) []int {
	if len(num1) == 0 {
		return num2
	}
	if len(num2) == 0 {
		return num1
	}

	minDigitNumber := num1
	maxDigitNumber := num2
	if len(num2) < len(num1) {
		minDigitNumber = num2
		maxDigitNumber = num1
	}

	result := make([]int, len(maxDigitNumber))
	j := len(minDigitNumber) - 1
	minDigit := minDigitNumber[j]
	var carry int
	for i := len(maxDigitNumber) - 1; i >= 0; i-- {
		sum := maxDigitNumber[i] + minDigit + carry
		result[i] += sum % 10
		carry = sum / 10

		if j >= 1 {
			j--
			minDigit = minDigitNumber[j]
		} else {
			minDigit = 0
		}
	}

	if carry > 0 {
		result = append([]int{1}, result...)
	}

	return result
}

// AddSliceOfTwoNumbers2 solves the problem in O(n) time and true O(1) extra space,
// by writing the sum directly into the longer of the two input slices instead of
// allocating a new result slice.
//
// This solution was provided by the AI.
func AddSliceOfTwoNumbers2(num1, num2 []int) []int {
	if len(num1) == 0 {
		return num2
	}
	if len(num2) == 0 {
		return num1
	}

	minDigitNumber := num1
	maxDigitNumber := num2
	if len(num2) < len(num1) {
		minDigitNumber = num2
		maxDigitNumber = num1
	}

	j := len(minDigitNumber) - 1
	carry := 0
	for i := len(maxDigitNumber) - 1; i >= 0; i-- {
		sum := maxDigitNumber[i] + carry
		if j >= 0 {
			sum += minDigitNumber[j]
			j--
		}
		maxDigitNumber[i] = sum % 10
		carry = sum / 10
	}

	if carry > 0 {
		maxDigitNumber = append([]int{carry}, maxDigitNumber...)
	}

	return maxDigitNumber
}

// AddSliceOfTwoNumbers3 solves the problem in O(n) time and O(1) space.
//
// This is the original solution in the repo.
func AddSliceOfTwoNumbers3(num1, num2 []int) []int {
	num1, num2 = equalizeLengths(num1, num2)
	carry := false
	for i := len(num1) - 1; i > -1; i-- {
		num1[i] += num2[i]
		if carry {
			num1[i]++
			carry = false
		}
		if num1[i] >= 10 {
			num1[i] -= 10
			carry = true
		}
	}
	if carry {
		num1 = append([]int{1}, num1...)
	}
	return num1
}

func equalizeLengths(num1, num2 []int) ([]int, []int) {
	diff := absDiff(len(num1), len(num2))
	zeros := make([]int, diff)
	if len(num2) > len(num1) {
		num1 = append(zeros, num1...)
	} else if len(num1) > len(num2) {
		num2 = append(zeros, num2...)
	}
	return num1, num2
}

func absDiff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
