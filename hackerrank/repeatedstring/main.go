package main

import "fmt"

func repeatedString(s string, n int64) int64 {
	var count int64
	var i int64
	var subTimes int64 = n/int64(len(s))
	var remainder int64 = int64(n % int64(len(s)))

	fmt.Println(remainder)
	for _, c := range s {
		if c == 'a' {
			count++
		}
	}
	count *= subTimes
	fmt.Println(count)

	for i = 0; i < remainder; i++ {
		if s[i] == 'a' {
			count++
		}
	}
	return count
}

func main() {
	s := "aba"
	var n int64 = 10
	fmt.Printf("count: %v\n", repeatedString(s, n))
}
