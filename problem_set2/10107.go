package main

import (
	"fmt"
	"sort"
)

type array []int

func isArrayEven(arr []int) bool {
	return len(arr) % 2 == 0
}

func (arr array) averageToMedian() int {
	a := len(arr) / 2
	return (arr[a - 1] + arr[a]) / 2
}

func (arr array) median() int {
	return arr[len(arr) / 2]
}

func main() {
	// Declare variables	
	var in, median int
	var nums array

	// Ask for input
	_, err := fmt.Scan(&in)

	// Append values to array
	for err == nil {
		nums = append(nums, in)

		// Sort the array
		sort.Ints(nums)

		// Check if array size is even
		// if even, assign the average of the two middle values to `median`
		// otherwise, assign the value which separates the array into two equal parts
		if isArrayEven(nums) {
			median = nums.averageToMedian()
		} else {
			median = nums.median()
		}

		// Display median
		println(median);

		// Ask for another input
		_, err = fmt.Scan(&in)
	}
}
