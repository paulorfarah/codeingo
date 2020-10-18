package main

import "fmt"

func subarraySum(arr []int, sum int) bool {
	i := 0
	bound := 0
	asum := arr[0]
	for bound < len(arr) {
		for asum > sum {
			asum -= arr[i]
			i++
		}
		if asum == sum {
			fmt.Printf("%v %v ", i+1, bound+1)
			return  true
		}
		fmt.Println(bound)
		bound++
		if bound < len(arr) {
			asum += arr[bound]
		}
	}
	return false
}


func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	res := subarraySum(arr, 15)
	fmt.Println(res)
}
