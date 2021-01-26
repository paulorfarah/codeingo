package main

import "fmt"

func freqQuery(queries [][]int32) []int32  {
	m := make(map[int32]int32)
	mr := make(map[int32]int32)
	r := make([]int32, 0)

	for _, v := range queries {
		fmt.Println(mr)
		if v[0] == 1 {
			if mr[m[v[1]]] > 0 {
				mr[m[v[1]]]--
			}
			m[v[1]]++
			mr[m[v[1]]]++
		} else if v[0] == 2 {
			mr[m[v[1]]]--
			if m[v[1]] > 0 {
				m[v[1]]--
			} else {
				m[v[1]] = 0
			}
			mr[m[v[1]]]++
		} else {
			if mr[v[1]] > 0 {
				r = append(r, 1)
			} else {
				r = append(r, 0)
			}
		}
	}
	return r
}

func main() {
	queries := [][]int32{{1, 1}, {1, 1}, {1, 1}, {2, 1}, {2,1}, {3, 1}}
 	fmt.Println(freqQuery(queries))
}
