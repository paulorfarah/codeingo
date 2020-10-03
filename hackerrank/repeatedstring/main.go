package main

import "fmt"

func repeatedString(s string, n int64) int64 {
	var (
		count int64
		i int64
		subTimes int64 = n/int64(len(s))
		remainder int64 = int64(n % int64(len(s)))
	)

	for _, c := range s {
		if c == 'a' {
			count++
		}
	}
	count *= subTimes

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
