package main

import "fmt"

func minimumSwaps(arr []int32) int32 {
	var swaps int32 
	min := 0
	minVal := arr[0]

	//find min value
	for i := 1; i < len(arr); i++ {
		if arr[i] < minVal {
			min = i
			minVal = arr[i]
		}
	}

	if min != 0 {
		aux := arr[0]
		arr[0] = arr[min]
		arr[min] = aux
		swaps++
	}

	for i := 1; i < len(arr) - 1; i++ {
		d := int(arr[i] - arr[0])
		for arr[d] != arr[i] {
			aux := arr[d]
			arr[d] = arr[i]
			arr[i] = aux
			swaps++
			d = int(arr[i] - arr[0])
		}
		i = d
	}
	return swaps
}

func main() {
	fmt.Println("teste")
	var arr = []int32{4, 3, 1, 2}
	res := minimumSwaps(arr)
	fmt.Println(res)
}
