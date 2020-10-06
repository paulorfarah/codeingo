package main

import "fmt"

func minimumBribes(q []int32) {
	bribes := 0
	for k, v := range q{
		if v-1 > int32(k+2){
			fmt.Println("Too chaotic")
			return
		}

		i := k-2
		if i < 0 {
			i = 0
		}
		for j:= i; j < k; j++ {
			if q[j] > v {
				bribes++
			}
		}
//		if v-1 > int32(k) {
//			bribes += int(v-1) - k
//		}
	}
	fmt.Println(bribes)
}

func main() {
	q := []int32{2, 1, 5, 3, 4}
	minimumBribes(q)
	q = []int32{2, 5, 1, 3, 4}
	minimumBribes(q)
	q = []int32{1, 2, 5, 3, 7, 8, 6, 4}
	minimumBribes(q)
}
