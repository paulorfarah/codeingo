package main

import (
	"fmt"
	"sort"
	"strings"
)


func SherlockAndAnagrams(s string) int {
	hashMap := make(map[string]int)
	r := []rune(s)
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			c := strings.Split(string(r[i:j+1]), "")
			sort.Strings(c)
			k := strings.Join(c, "")
			if _, ok := hashMap[k]; ok {
				hashMap[k] += 1
			} else {
				hashMap[k] = 1
			}
		}
	}
	pairs := 0;
	for _, v := range hashMap {
		pairs += (v * (v - 1)) / 2
	}
	return pairs
}



func main() {
	s := ""
	q := SherlockAndAnagrams(s)
	fmt.Println(q)

}
