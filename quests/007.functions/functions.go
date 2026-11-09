package functions

import "fmt"

// Divide returns a / b or an error if b == 0
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero is not allowed")
	}
	return a / b, nil
}

// SumAll returns the sum of all provided integers
func SumAll(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}
	sum := 0
	for _, num := range nums {
		sum += num
	}

	return sum
}

// MaxMin returns the max and min of all provided integers
// Returns an error if no numbers are provided
func MaxMin(nums ...int) (int, int, error) {
	if len(nums) == 0 {
		return 0, 0, fmt.Errorf("no numbers provided")
	}
	maxi := -9999
	mini := 9999
	for _, values := range nums {
		if values > maxi {
			maxi = values
		}
		if values < mini {
			mini = values
		}
	}
	return maxi, mini, nil
}

// ConcatAll joins all strings using the provided separator
func ConcatAll(sep string, strs ...string) string {
	if len(strs) == 0 {
		return ""
	}
	res := ""
	for i, str := range strs {
		res += str
		if i < len(strs)-1 {

			res += sep
		}
	}
	return res
}
