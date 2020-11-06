package main

import "fmt"

func checkMagazine(magazine []string, note []string) {
	magMap := make(map[string]bool)
	for _, word := range magazine {
		magMap[word] = true
	}
	var i = true
	for _, word := range note {
		_, ok := magMap[word]
		if ok == false {
			i = false
			break
		}else {
			delete(magMap, word)
		}
	}
	if i == true {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func main() {
	fmt.Println("1")
	mag := []string{"give", "me", "one", "grand", "today", "night"}
	note := []string{"give", "one", "grand", "today"}
	checkMagazine(mag, note)

	fmt.Println("2")
	mag = []string{"two", "times", "three", "is", "not", "four"}
	note = []string{"two", "times", "two", "is", "four"}
	checkMagazine(mag, note)

	fmt.Println("3")
	mag = []string{"ive", "got", "a", "lovely", "bunch", "of", "coconuts"}
	note = []string{"ive", "got", "some", "coconuts"}
	checkMagazine(mag, note)
}


