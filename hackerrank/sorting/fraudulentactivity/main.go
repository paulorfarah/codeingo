package main

import "fmt"

func activityNotifications(expenditure []int32, d int32) int32 {
	n := int32(0)
	var median int32
	aux := d % 2
	for i := d; i < int32(len(expenditure)); i++ {
		qs(expenditure, i - d, i - 1)
		if aux == 0 {
			median = (expenditure[(i - d) + d/2] + expenditure[(i - d) + d/2 + 1])/2
		} else {
			median = expenditure[(i - d) + d/2]
		}
		if expenditure[i] >= 2 * median {
			n++
		}
	}
	return n
}

func qs(a []int32, l int32, r int32) {
	if l > r {
		return
	}
	p := partition(a, l, r)
	qs(a, l, p-1)
	qs(a, p+1, r)
}

func partition(a []int32, l, r int32) int32 {
	p := a[r]
	i := l - 1
	for j := l; j < r; j++ {
		if a[j] < p {
			i++
			tmp := a[i]
			a[i] = a[j]
			a[j] = tmp
		}
	}
	i++
	tmp := a[i]
	a[i] = p
	a[r] = tmp
	return i
}

func main() {
	a := []int32{2, 3, 4, 2, 3, 6, 8, 4, 5}
	fmt.Println(activityNotifications(a, 5))
}
