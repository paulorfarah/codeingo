package main

import "fmt"

func countTriplets(arr []int64, r int64) {
	count := 0
	for i:=0; i < len(arr) -  2; i++ {
		for j := 1; j < len(arr) - 1 ; j++ {
			for k := 2; k< len(arr); k++ {
				fmt.Println("i: ", arr[i], " j: ", arr[j], " k: ", arr[k])
				if arr[i]*r == arr[j] && arr[j]*r == arr[k] {
					count++
				}
			}
		}
	}
	fmt.Println(count)
}

func main() {
	arr := []int64{1, 3, 9, 9, 27, 81}
	countTriplets(arr, 3)
}
