package main

import "fmt"

func main() {
	s := "cdcd"
	var list []string
	r := []rune(s)
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			list = append(list, string(r[i:j]))
		}
	}

	q := 0
	for i := 0; i < len(list) - 1; i++ {
		r1 := []rune(list[i])
		for k := i + 1; k < len(list); k++ {
			m := make(map[string]int)
			for j:= 0; j < len(r1); j++ {
				m[string(r1[j])]++
			}
		//	fmt.Println(m)
			r2 := []rune(list[k])
			fmt.Println(string(r1), string(r2))
			for j := 0; j < len(r2); j++ {
				m[string(r2[j])]--
			}
			//fmt.Println(m)

			equal := true
			for _, v := range m {
				if v != 0 {
					equal = false
				}
			}
			if equal {
				q++
			}
		}
	}
	fmt.Println(q)

}
