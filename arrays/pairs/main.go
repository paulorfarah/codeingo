package main

import "fmt"

func CountPairs(data []int) int {
	m := make(map[int]int)
	for _, v := range data {
		m[v]++
	}
	cont := 0
	for _, v := range m {
		cont += v/2
	}
	return cont
}

func main(){
	numbers := []int{1,2,3,2,1,3,1,2,3,2,1,1,1}
	fmt.Println(CountPairs(numbers))
}
