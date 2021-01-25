package main

import "fmt"

func countTriplets(arr []int64, r int64) {
	count := int64(0)
	v2 := make(map[int64]int64)
	v3 := make(map[int64]int64)
	for _, v := range arr {
		count += v3[v]
		v3[v*r] += v2[v]
		v2[v*r] += 1
	}
	fmt.Println(count)
}

func main() {
	arr := []int64{1,2,2,4} //1, 3, 9, 9, 27, 81}
	countTriplets(arr, 2)
}
