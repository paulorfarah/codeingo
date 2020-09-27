package main

import (
	"fmt"
)

func ReverseStrings(s []string) []string {
	first := 0
	last := len(s) - 1
	for first < last {
		s[first], s[last] = s[last], s[first]
		first++
		last--
	}
	return s
}

func main() {
	s := []string{"a", "b", "c"}
	fmt.Printf("Original: %v\n", s)
	fmt.Printf("Reversed: %v\n", ReverseStrings(s))
}
