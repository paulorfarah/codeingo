package main

import (
	"fmt"
)

func main(){
	w1 := "hello"
	w2 := "world"

	d1 := make(map[string]bool)
	for _, i := range(w1){ 
		fmt.Println(i)
		d1[string(i)] = true
	}
	for _, i := range w2{
		_, found := d1[string(i)]
		if found {
			//return "YES"
			fmt.Println("YES")
			return
		}
	}
//	return "NO"
	fmt.Println("NO")
}
