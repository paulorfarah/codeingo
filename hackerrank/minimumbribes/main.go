package main

import "fmt"


func minimumBribes(q []int32) {
	chaotic := false
	bribes := 0
	for k, v := range q {
		if v - int32(( k + 1 )) > int32(2) {
			chaotic = true
		}
		for j := int(v) - 2; j < k; j++ {
			if j < 0 {
				j = 0
			}
			if q[j] > v {
				bribes++
			}
		}
	}
	if chaotic == true {
		fmt.Println("too chaotic")
	} else {
		fmt.Println(bribes)
	}
}
func main() {
 	x := []int32{2, 1, 5, 4, 3}
	minimumBribes(x)
}
