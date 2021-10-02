package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := "ab-cd"
	r := reverseOnlyLetters(s)
	fmt.Println(r)
}

func reverseOnlyLetters(str string) string {
	s := []rune(str)
	i := 0
	j := len(s)-1
	for i < j {
		for i < j && !unicode.IsLetter(s[i]) {
			i++
		}
		for j > i && !unicode.IsLetter(s[j]) {
			j--
		}

		temp := s[i]
		s[i]=s[j]
		i++
		s[j]=temp
		j--
	}
	return string(s)
}

