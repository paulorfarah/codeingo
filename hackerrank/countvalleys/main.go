package main 

import "fmt"

func countValleys(steps int32, path string) int32 {
	count := 0
	var valleys int32 
	typeH := 's'
	for _, c := range path {
	    if c == 'D' {
		    count--
	    } else {
		    count ++
	    }
	    if typeH == 's' {
		    if count < 0 {
			    typeH = 'v'
		    }else if count > 0 {
			    typeH = 'm'
		    }
	    }else if count == 0 {
		    if typeH == 'v' {
			    valleys++
		    }
		    typeH = 's'
	    }
    	}
	return valleys
}

func main() {
	path := "DDUUUUDD"
	v := countValleys(8, path)
	fmt.Println(v)

	//test case 2
	path = "UDDDUDUU"
	v = countValleys(8, path)
	fmt.Println(v)

	//test case 3
	path = "UUUUUUUU"
	v = countValleys(8, path)
	fmt.Println(v)
	
	//test case 4
	path = "DDDDDDDD"
	v = countValleys(8, path)
	fmt.Println(v)

	//test case 5
	path = "UDDUUDDU"
	v = countValleys(8, path)
	fmt.Println(v)

}
