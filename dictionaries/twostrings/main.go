package main

import (
	"fmt"
)

func main(){
	w1 := "hello"
	w2 := "world"

	d1 := make(map[string]bool)
	r1 := []rune(w1)
	r2 := []rune(w2)
	for i := 0; i <= len(w1); i++ {
		for j := i+1; j <= len(w1); j++ {
			fmt.Println(string(r1[i:j]))
			d1[string(r1[i:j])] = true
		}
	}
	for i := 0; i <= len(w2); i++ {
		for j := i+1; j <= len(w2); j++ {
			_, found := d1[string(r2[i:j])]
			if found {
				//return "YES"
				fmt.Println("YES")
			}
		}
	}
//	return "NO"
	fmt.Println("NO")
}
