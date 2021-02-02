package main

import "fmt"

func main() {
	a := [5]int{2, 9, 5, 7, 1}
	fmt.Println(len(a))
	for i := len(a) -1; i > 0; i-- {
		for j := 0; j < i; j++ {
			fmt.Println(a[j], ", ", a[j+1])
			if a[j] > a[j+1] {
				tmp := a[j+1]
				a[j+1] = a[j]
				a[j] = tmp
			}
		}
	}
	fmt.Println(a)
}
