package main

import "fmt"

func hourglass(arr [][]int32) int32 {
	var (
		i int32
		j int32
	)
	max := int32(-2147483648)

	for i=int32(0); i < int32(len(arr)-2); i++ {
		for j=int32(0); j < int32(len(arr[0])-2); j++ {
			sum := arr[i][j] + arr[i][j+1] + arr[i][j+2] + arr[i+1][j+1] + arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]
			if sum > max {
				max = sum
			}
		}
	}
	return max
}

func main() {
	a := [][]int32{
		{1,1,1,0,0,0},
		{0,1,0,0,0,0},
		{1,1,1,0,0,0},
		{0,0,2,4,4,0},
		{0,0,0,2,0,0},
		{0,0,1,2,4,0},
	}
	fmt.Printf("max: %v", hourglass(a))
}
