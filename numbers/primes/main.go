package main

import (
	"fmt"
)

func PrimeNumbers(n int) []int {
	numbers := make([]bool, n+1)
	for i:=2; i<=n; i++ {
		numbers[i] = true
	}

	for p:=2; p*p<=n; p++{
		if numbers[p] == true {
			for i:=p*2; i <= n; i+=p {
				fmt.Println(i)
				numbers[i] = false
			}
		}
	}
	var primes []int
	for i:=2; i<=n; i++ {
		if numbers[i] == true {
			primes = append(primes, i)
		}
	}
	return primes
}

func main() {
	primes := PrimeNumbers(10)
	for _, i := range primes{
		fmt.Printf("%v ", i)
	}
}
