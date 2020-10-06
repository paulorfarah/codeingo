package main

import "fmt"

func qs(arr []int, l int, r int) {
	if l >= r {
		return
	}
	p := partition(arr, l, r)
	qs(arr, l, p - 1)
	qs(arr, p+1, r)
}

func partition(array []int, l int, r int) {

}
