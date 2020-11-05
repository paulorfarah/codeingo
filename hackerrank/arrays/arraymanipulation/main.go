package main

import "fmt"

func arrayManipulation(n int32, queries [][]int32) int64 {
	a := make([]int64, n+1)
	sum := int64(0)
	max := int64(0)

	for _, v := range queries {
		a[v[0]-1] += int64(v[2])
		a[v[1]] -= int64(v[2])
	}
	for i:=0; i < int(n); i++ {
		sum += a[i]
		if sum > max {
			max = sum
		}
	}
	return max
}

func main() {
	q := [][]int32{
		{1, 5, 3},
		{4, 8, 7},
		{6, 9, 1},
	}
	n := 10
	max := arrayManipulation(int32(n), q)
	fmt.Println(max)
	q = [][]int32{
		{1, 2, 100},
		{2, 5, 100},
		{3, 4, 100},
	}
	n = 5
	max = arrayManipulation(int32(n), q)
	fmt.Println(max)

}
