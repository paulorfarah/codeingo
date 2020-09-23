package main

import (
	"fmt"
)

func BinarySearch(key int, numbers []int) int {
	min := 0
	max := len(numbers)
	for min <= max {
		mid := (min + max) / 2
		if key == numbers[mid] {
			return mid
		} else if key < numbers[mid] {
			max = mid - 1
		} else {
			min = mid + 1
		}
	}
	return -1
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	pos := BinarySearch(2, numbers)
	fmt.Println("pos", pos)
}
