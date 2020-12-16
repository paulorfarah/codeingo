package main

import "fmt"


func main() {
	magazine := [6]string{"two", "times", "three", "is", "not", "four"}
	note := [5]string{"two", "times", "two", "is", "four"}
	
	magD := make(map[string]int)
	for _, m := range magazine {
		magD[m]++
	}
//	res := "Yes"
	for _, n := range note {
		fmt.Println(n, magD[n])
		magD[n]--

	}
}
