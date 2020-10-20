package main

import "fmt"


func minimumBribes(q []int32) {
	chaotic := false
	bribes := 0
	for i := 0; i < len(q); i++ {
		if q[i] - int32((i+1)) > int32(2) {
			chaotic = true
		}
		for j := int(q[i]) - 2; j < i; j++ {
			if j < 0 {
				j = 0
			}
			if q[j] > q[i] {
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
